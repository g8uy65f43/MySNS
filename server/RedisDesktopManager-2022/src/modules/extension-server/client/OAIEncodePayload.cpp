/**
 * RESP.app Extension server
 * RESP.app Extension Server API allows you to extend RESP.app with your custom data formatters
 *
 * The version of the OpenAPI document: 2022.0-preview1
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

#include "OAIEncodePayload.h"

#include <QDebug>
#include <QJsonArray>
#include <QJsonDocument>
#include <QObject>

#include "OAIHelpers.h"

namespace RespExtServer {

OAIEncodePayload::OAIEncodePayload(QString json) {
    this->initializeModel();
    this->fromJson(json);
}

OAIEncodePayload::OAIEncodePayload() {
    this->initializeModel();
}

OAIEncodePayload::~OAIEncodePayload() {}

void OAIEncodePayload::initializeModel() {

    m_data_isSet = false;
    m_data_isValid = false;

    m_metadata_isSet = false;
    m_metadata_isValid = false;
}

void OAIEncodePayload::fromJson(QString jsonString) {
    QByteArray array(jsonString.toStdString().c_str());
    QJsonDocument doc = QJsonDocument::fromJson(array);
    QJsonObject jsonObject = doc.object();
    this->fromJsonObject(jsonObject);
}

void OAIEncodePayload::fromJsonObject(QJsonObject json) {

    m_data_isValid = ::RespExtServer::fromJsonValue(data, json[QString("data")]);
    m_data_isSet = !json[QString("data")].isNull() && m_data_isValid;

    m_metadata_isValid = ::RespExtServer::fromJsonValue(metadata, json[QString("metadata")]);
    m_metadata_isSet = !json[QString("metadata")].isNull() && m_metadata_isValid;
}

QString OAIEncodePayload::asJson() const {
    QJsonObject obj = this->asJsonObject();
    QJsonDocument doc(obj);
    QByteArray bytes = doc.toJson();
    return QString(bytes);
}

QJsonObject OAIEncodePayload::asJsonObject() const {
    QJsonObject obj;
    if (m_data_isSet) {
        obj.insert(QString("data"), ::RespExtServer::toJsonValue(data));
    }
    if (m_metadata_isSet) {
        obj.insert(QString("metadata"), ::RespExtServer::toJsonValue(metadata));
    }
    return obj;
}

QString OAIEncodePayload::getData() const {
    return data;
}
void OAIEncodePayload::setData(const QString &data) {
    this->data = data;
    this->m_data_isSet = true;
}

bool OAIEncodePayload::is_data_Set() const{
    return m_data_isSet;
}

bool OAIEncodePayload::is_data_Valid() const{
    return m_data_isValid;
}

OAIObject OAIEncodePayload::getMetadata() const {
    return metadata;
}
void OAIEncodePayload::setMetadata(const OAIObject &metadata) {
    this->metadata = metadata;
    this->m_metadata_isSet = true;
}

bool OAIEncodePayload::is_metadata_Set() const{
    return m_metadata_isSet;
}

bool OAIEncodePayload::is_metadata_Valid() const{
    return m_metadata_isValid;
}

bool OAIEncodePayload::isSet() const {
    bool isObjectUpdated = false;
    do {
        if (m_data_isSet) {
            isObjectUpdated = true;
            break;
        }

        if (m_metadata_isSet) {
            isObjectUpdated = true;
            break;
        }
    } while (false);
    return isObjectUpdated;
}

bool OAIEncodePayload::isValid() const {
    // only required properties are required for the object to be considered valid
    return true;
}

} // namespace RespExtServer
