QT       += core gui network concurrent widgets quick quickwidgets testlib

TARGET = tests
TEMPLATE = app

CONFIG += debug c++17
CONFIG-=app_bundle 

PROJECT_ROOT = $$PWD/../..//
SRC_DIR = $$PROJECT_ROOT/src//

HEADERS += \
    $$PROJECT_ROOT/3rdparty/qredisclient/tests/unit_tests/basetestcase.h \    
    $$files($$PROJECT_ROOT/3rdparty/qredisclient/tests/unit_tests/mocks/*.h) \
    $$files($$PROJECT_ROOT/src/modules/common/*.h) \
    $$files($$PWD/*.h) \

SOURCES += \
    $$PROJECT_ROOT/3rdparty/qredisclient/tests/unit_tests/basetestcase.cpp \
    $$files($$PROJECT_ROOT/3rdparty/qredisclient/tests/unit_tests/mocks/*.cpp) \
    $$files($$PROJECT_ROOT/src/modules/common/*.cpp) \
    $$PWD/main.cpp \

INCLUDEPATH += $$SRC_DIR/modules/ \
    $$SRC_DIR/ \
    $$PWD/ \
    $$PROJECT_ROOT/3rdparty/qredisclient/tests/unit_tests/ \
    $$PROJECT_ROOT/3rdparty/fakeit/single_header/qtest/

DEFINES += INTEGRATION_TESTS

#TEST CASES
include($$PWD/testcases/app/app-tests.pri)
include($$PWD/testcases/connections-tree/connections-tree-tests.pri)
include($$PWD/testcases/console/console-tests.pri)
include($$PWD/testcases/value-editor/value-editor-tests.pri)
#############
include($$PROJECT_ROOT/3rdparty/3rdparty.pri)

release: DESTDIR = $$PROJECT_ROOT/bin/tests
debug:   DESTDIR = $$PROJECT_ROOT/bin/tests

UI_DIR = $$DESTDIR/ui
OBJECTS_DIR = $$DESTDIR/obj
MOC_DIR = $$DESTDIR/obj
RCC_DIR = $$DESTDIR/obj
