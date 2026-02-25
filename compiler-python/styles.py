"""Application stylesheets."""


MAIN_STYLE = """
* {
    font-family: "Segoe UI", "Arial", sans-serif;
    color: #333333;
}

QMainWindow {
    background-color: #FFFFFF;
    border: none;
}

QStatusBar {
    background-color: #F0F0F0;
    border-top: 1px solid #CCCCCC;
    color: #666666;
}

QSplitter::handle {
    background-color: #E0E0E0;
    width: 3px;
}

QSplitter::handle:hover {
    background-color: #D0D0D0;
}

QTextEdit, QPlainTextEdit {
    border: 1px solid #CCCCCC;
    border-radius: 3px;
    padding: 4px;
    background-color: #FFFFFF;
}

QTextEdit:focus, QPlainTextEdit:focus {
    border: 2px solid #0078D4;
}
"""

TOOLBAR_STYLE = """
QToolBar {
    border: none;
    spacing: 0px;
    padding: 4px;
    background-color: #F5F5F5;
}

QToolBar QToolButton {
    padding: 3px;
    border: 1px solid transparent;
    border-radius: 4px;
    background-color: transparent;
    color: #000000;
    icon-size: 24px;
}

QToolBar QToolButton:hover {
    background-color: #E0E0E0;
    border: 1px solid #C0C0C0;
}

QToolBar QToolButton:pressed {
    background-color: #C0C0C0;
}

QToolBar::separator {
    background-color: #D0D0D0;
    width: 1px;
    margin: 4px;
}
"""
