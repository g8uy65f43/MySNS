import QtQuick 2.13
import QtQuick.Layouts 1.1
import "./../../common"

RowLayout {

    BetterLabel {
        text: qsTranslate("RESP", "Order of elements:")
    }

    BetterComboBox {
        id: filterDirection

        enabled: !keyTab.keyModel.singlePageMode

        ListModel {
            id: filterDirectionModel

            Component.onCompleted: {
                filterDirectionModel.append({ value: "default", text: qsTranslate("RESP", "Default") })
                filterDirectionModel.append({ value: "reverse", text: qsTranslate("RESP", "Reverse") })
                filterDirection.currentIndex = 0
            }
        }

        textRole: "text"

        model: filterDirectionModel
        onCurrentIndexChanged: {
            var direction = filterDirectionModel.get(currentIndex)["value"];
            keyModel.setFilter("order", direction);
            reloadValue();
        }
    }

    Item { Layout.fillWidth: true }
}
