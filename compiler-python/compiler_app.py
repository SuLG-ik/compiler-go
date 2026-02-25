from __future__ import annotations

import os
from PySide6.QtWidgets import (
    QMainWindow, QTextEdit, QPlainTextEdit, QSplitter,
    QToolBar, QStatusBar, QStyle
)
from PySide6.QtGui import QAction, QKeySequence, QFont, QIcon
from PySide6.QtCore import Qt, QSize

from strings import Strings
from styles import TOOLBAR_STYLE


class CompilerWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self._current_file: str | None = None
        self._init_ui()
        self._update_title()

    def _init_ui(self):
        self.setMinimumSize(900, 600)
        self.resize(1100, 700)

        splitter = QSplitter(Qt.Orientation.Vertical)
        self.setCentralWidget(splitter)

        self.editor = QPlainTextEdit()
        self.editor.setReadOnly(False)
        self.editor.setFont(QFont("Courier New", 11))
        self.editor.setTabStopDistance(32)
        self.editor.setPlaceholderText("Код здесь...")
        self.editor.setLineWrapMode(QPlainTextEdit.LineWrapMode.NoWrap)
        splitter.addWidget(self.editor)

        self.output = QTextEdit()
        self.output.setReadOnly(True)
        self.output.setFont(QFont("Courier New", 11))
        self.output.setStyleSheet("background-color: #F5F5F5;")
        self.output.setPlaceholderText(Strings.OUTPUT_PLACEHOLDER)
        splitter.addWidget(self.output)

        splitter.setStretchFactor(0, 3)
        splitter.setStretchFactor(1, 1)

        self.setStatusBar(QStatusBar())

        self._create_actions()
        self._create_menus()
        self._create_toolbar()

    def _create_actions(self):
        style = self.style()

        self.act_new = QAction(style.standardIcon(QStyle.StandardPixmap.SP_FileIcon), Strings.Actions.NEW, self)
        self.act_new.setShortcut(QKeySequence.StandardKey.New)
        self.act_new.setToolTip(Strings.Tooltips.NEW)

        self.act_open = QAction(style.standardIcon(QStyle.StandardPixmap.SP_DialogOpenButton), Strings.Actions.OPEN, self)
        self.act_open.setShortcut(QKeySequence.StandardKey.Open)
        self.act_open.setToolTip(Strings.Tooltips.OPEN)

        self.act_save = QAction(style.standardIcon(QStyle.StandardPixmap.SP_DialogSaveButton), Strings.Actions.SAVE, self)
        self.act_save.setShortcut(QKeySequence.StandardKey.Save)
        self.act_save.setToolTip(Strings.Tooltips.SAVE)

        self.act_save_as = QAction(Strings.Actions.SAVE_AS, self)
        self.act_save_as.setShortcut(QKeySequence("Ctrl+Shift+S"))
        self.act_save_as.setToolTip(Strings.Tooltips.SAVE_AS)

        self.act_exit = QAction(Strings.Actions.EXIT, self)
        self.act_exit.setShortcut(QKeySequence("Alt+F4"))
        self.act_exit.setToolTip(Strings.Tooltips.EXIT)

        self.act_undo = QAction(style.standardIcon(QStyle.StandardPixmap.SP_ArrowBack), Strings.Actions.UNDO, self)
        self.act_undo.setShortcut(QKeySequence.StandardKey.Undo)
        self.act_undo.setToolTip(Strings.Tooltips.UNDO)

        self.act_redo = QAction(style.standardIcon(QStyle.StandardPixmap.SP_ArrowForward), Strings.Actions.REDO, self)
        self.act_redo.setShortcut(QKeySequence.StandardKey.Redo)
        self.act_redo.setToolTip(Strings.Tooltips.REDO)

        self.act_cut = QAction(QIcon.fromTheme("edit-cut"), Strings.Actions.CUT, self)
        self.act_cut.setShortcut(QKeySequence.StandardKey.Cut)
        self.act_cut.setToolTip(Strings.Tooltips.CUT)

        self.act_copy = QAction(QIcon.fromTheme("edit-copy"), Strings.Actions.COPY, self)
        self.act_copy.setShortcut(QKeySequence.StandardKey.Copy)
        self.act_copy.setToolTip(Strings.Tooltips.COPY)

        self.act_paste = QAction(QIcon.fromTheme("edit-paste"), Strings.Actions.PASTE, self)
        self.act_paste.setShortcut(QKeySequence.StandardKey.Paste)
        self.act_paste.setToolTip(Strings.Tooltips.PASTE)

        self.act_delete = QAction(Strings.Actions.DELETE, self)
        self.act_delete.setShortcut(QKeySequence.StandardKey.Delete)
        self.act_delete.setToolTip(Strings.Tooltips.DELETE)

        self.act_select_all = QAction(Strings.Actions.SELECT_ALL, self)
        self.act_select_all.setShortcut(QKeySequence.StandardKey.SelectAll)
        self.act_select_all.setToolTip(Strings.Tooltips.SELECT_ALL)

        self.act_task = QAction(Strings.Actions.TASK, self)
        self.act_grammar = QAction(Strings.Actions.GRAMMAR, self)
        self.act_classification = QAction(Strings.Actions.CLASSIFICATION, self)
        self.act_method = QAction(Strings.Actions.METHOD, self)
        self.act_test_example = QAction(Strings.Actions.TEST_EXAMPLE, self)
        self.act_references = QAction(Strings.Actions.REFERENCES, self)
        self.act_source_code = QAction(Strings.Actions.SOURCE_CODE, self)

        self.act_run = QAction(style.standardIcon(QStyle.StandardPixmap.SP_MediaPlay), Strings.Actions.RUN, self)
        self.act_run.setShortcut(QKeySequence("Ctrl+R"))
        self.act_run.setToolTip(Strings.Tooltips.RUN)

        self.act_help = QAction(style.standardIcon(QStyle.StandardPixmap.SP_DialogHelpButton), Strings.Actions.HELP, self)
        self.act_help.setShortcut(QKeySequence.StandardKey.HelpContents)
        self.act_help.setToolTip(Strings.Tooltips.HELP)

        self.act_about = QAction(style.standardIcon(QStyle.StandardPixmap.SP_MessageBoxInformation), Strings.Actions.ABOUT, self)
        self.act_about.setToolTip(Strings.Tooltips.ABOUT)

    def _create_menus(self):
        menubar = self.menuBar()

        file_menu = menubar.addMenu(Strings.Menus.FILE)
        file_menu.addAction(self.act_new)
        file_menu.addAction(self.act_open)
        file_menu.addAction(self.act_save)
        file_menu.addAction(self.act_save_as)
        file_menu.addSeparator()
        file_menu.addAction(self.act_exit)

        edit_menu = menubar.addMenu(Strings.Menus.EDIT)
        edit_menu.addAction(self.act_undo)
        edit_menu.addAction(self.act_redo)
        edit_menu.addSeparator()
        edit_menu.addAction(self.act_cut)
        edit_menu.addAction(self.act_copy)
        edit_menu.addAction(self.act_paste)
        edit_menu.addAction(self.act_delete)
        edit_menu.addSeparator()
        edit_menu.addAction(self.act_select_all)

        text_menu = menubar.addMenu(Strings.Menus.TEXT)
        text_menu.addAction(self.act_task)
        text_menu.addAction(self.act_grammar)
        text_menu.addAction(self.act_classification)
        text_menu.addAction(self.act_method)
        text_menu.addAction(self.act_test_example)
        text_menu.addAction(self.act_references)
        text_menu.addAction(self.act_source_code)

        menubar.addAction(self.act_run)

        help_menu = menubar.addMenu(Strings.Menus.HELP)
        help_menu.addAction(self.act_help)
        help_menu.addAction(self.act_about)

    def _create_toolbar(self):
        toolbar = QToolBar(Strings.Toolbar.NAME)
        toolbar.setIconSize(QSize(24, 24))
        toolbar.setMovable(False)
        toolbar.setStyleSheet(TOOLBAR_STYLE)
        self.addToolBar(toolbar)

        toolbar.addAction(self.act_new)
        toolbar.addAction(self.act_open)
        toolbar.addAction(self.act_save)
        toolbar.addSeparator()
        toolbar.addAction(self.act_undo)
        toolbar.addAction(self.act_redo)
        toolbar.addAction(self.act_copy)
        toolbar.addAction(self.act_cut)
        toolbar.addAction(self.act_paste)
        toolbar.addAction(self.act_run)
        toolbar.addAction(self.act_help)
        toolbar.addAction(self.act_about)

    def _update_title(self):
        name = os.path.basename(self._current_file) if self._current_file else Strings.UNTITLED
        self.setWindowTitle(f"{name} — {Strings.APP_TITLE}")
