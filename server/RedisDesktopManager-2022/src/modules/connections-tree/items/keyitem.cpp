#include "keyitem.h"
#include <qredisclient/utils/text.h>
#include <QCoreApplication>
#include <QMenu>
#include <QMessageBox>
#include <QSettings>

#include "app/apputils.h"
#include "connections-tree/items/abstractnamespaceitem.h"
#include "connections-tree/model.h"
#include "connections-tree/utils.h"

using namespace ConnectionsTree;

QSharedPointer<AbstractNamespaceItem> parentTreeItemToNs(
    QWeakPointer<TreeItem> p) {
  auto parentNs = p.toStrongRef();

  if (!parentNs || !parentNs.staticCast<AbstractNamespaceItem>())
    return QSharedPointer<AbstractNamespaceItem>();

  return parentNs.staticCast<AbstractNamespaceItem>();
}

KeyItem::KeyItem(const QByteArray& fullPath, QWeakPointer<TreeItem> parent,
                 Model& model, bool shortNameRendering)
    : TreeItem(model),
      m_fullPath(fullPath),
      m_parent(parent),
      m_removed(false),
      m_shortRendering(shortNameRendering)
{
}

QString KeyItem::getDisplayName() const {
  QString title;

  if (m_parent && m_parent.toStrongRef()->type() == "namespace" &&
      m_shortRendering) {
    auto parent = parentTreeItemToNs(m_parent);

    title = printableString(getFullPath().mid(
        parent->getFullPath().size() +
        parent->operations()->getNamespaceSeparator().size()));
  } else {
    title = printableString(getFullPath(), true);
  }

  if (m_usedMemory > 0) {
    title.append(QString(" <b>[%1]</b>").arg(humanReadableSize(m_usedMemory)));
  }

  return title;
}

QByteArray KeyItem::getName() const { return getFullPath(); }

QList<QSharedPointer<TreeItem>> KeyItem::getAllChilds() const {
  return QList<QSharedPointer<TreeItem>>();
}

bool KeyItem::supportChildItems() const { return false; }

uint KeyItem::childCount(bool) const { return 0u; }

QSharedPointer<TreeItem> KeyItem::child(uint) {
  return QSharedPointer<TreeItem>();
}

QWeakPointer<TreeItem> KeyItem::parent() const { return m_parent; }

bool KeyItem::isEnabled() const {
  if (!m_removed && m_parent) {
    return m_parent.toStrongRef()->isEnabled();
  } else {
    return m_removed == false;
  }
}

QByteArray KeyItem::getFullPath() const { return m_fullPath; }

int KeyItem::getDbIndex() const {
  auto parentNs = parentTreeItemToNs(m_parent);

  if (!parentNs) {
    return -1;
  }

  return parentNs->getDbIndex();
}

void KeyItem::setRemoved() {
  m_removed = true;

  emit m_model.itemChanged(getSelf());
}

void KeyItem::getMemoryUsage(std::function<void(qlonglong)> callback) {
  auto parentNs = parentTreeItemToNs(m_parent);

  if (!parentNs || !parentNs->operations()) return callback(0);

  auto cb = QSharedPointer<Operations::GetUsedMemoryCallback>(
      new Operations::GetUsedMemoryCallback(
          getSelf(), [this, callback](qlonglong result) {
            m_usedMemory = result;
            callback(result);
            emit m_model.itemChanged(getSelf());
          }));

  parentNs->operations()->getUsedMemory(
      {getFullPath()}, getDbIndex(), cb,
      QSharedPointer<Operations::GetUsedMemoryCallback>());
}

void KeyItem::setFullPath(const QByteArray& p) {
  m_fullPath = p;

  emit m_model.itemChanged(getSelf());
}

QHash<QString, std::function<bool()>> KeyItem::eventHandlers() {
  auto events = TreeItem::eventHandlers();

  events.insert("click", [this]() {
    if (!isEnabled()) return true;

    auto parentNs = parentTreeItemToNs(m_parent);

    if (!parentNs || !parentNs->operations()) return true;

    parentNs->operations()->openKeyTab(
        getSelf().toStrongRef().staticCast<KeyItem>(), false);
    return true;
  });

  events.insert("mid-click", [this]() {
    if (!isEnabled()) return true;

    auto parentNs = parentTreeItemToNs(m_parent);

    if (!parentNs || !parentNs->operations()) return true;

    parentNs->operations()->openKeyTab(
        getSelf().toStrongRef().staticCast<KeyItem>(), true);

    return true;
  });

  events.insert("delete", [this]() {
    confirmAction(
        nullptr,
        QCoreApplication::translate("RESP",
                                    "Do you really want to delete this key?"),
        [this]() {
          auto parentNs = parentTreeItemToNs(m_parent);

          if (!parentNs || !parentNs->operations()) return;

          auto callback = QSharedPointer<Operations::DeleteDbKeyCallback>(
              new Operations::DeleteDbKeyCallback(
                  getSelf(), [this](const QString& err) {
                    emit m_model.error(QCoreApplication::translate(
                                           "RESP", "Cannot delete key:\n\n") +
                                       err);
                    return;
                  }));

          parentNs->operations()->deleteDbKey(*this, callback);
        });
    return true;
  });
  return events;
}
