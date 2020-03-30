#include "TextEditorWrapper.h"
#include "TextEditor.h"
#include "WrapperConverter.h"

IggTextEditor IggNewTextEditor()
{
  TextEditor *editor = new TextEditor();
  return static_cast<IggTextEditor>(editor);
}

void IggTextEditorRender(IggTextEditor handle, const char* aTitle, IggVec2 const *size, int aBorder)
{
  Vec2Wrapper sizeArg(size);
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->Render(aTitle, *sizeArg, aBorder != 0);
}

void IggTextEditorSetShowWhitespaces(IggTextEditor handle, int aValue)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetShowWhitespaces(aValue != 0);
}

void IggTextEditorSetTabSize(IggTextEditor handle, int size)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetTabSize(size);
}

void IggTextEditorSetText(IggTextEditor handle, const char* text)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetText(text);
}

IggBool IggTextEditorHasSelection(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  return editor->HasSelection() ? 1 : 0; 
}

char const *IggTextEditorGetText(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  std::string str = editor->GetText();
  char* c = strcpy(new char[str.length() + 1], str.c_str());
  return c;
}

char const *IggTextEditorGetSelectedText(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  std::string str = editor->GetSelectedText();
  char* c = strcpy(new char[str.length() + 1], str.c_str());
  return c;
}

char const *IggTextEditorGetCurrentLineText(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  std::string str = editor->GetCurrentLineText();
  char* c = strcpy(new char[str.length() + 1], str.c_str());
  return c;
}

IggBool IggTextEditorIsTextChanged(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  return editor->IsTextChanged() ? 1 : 0;
}

void IggTextEditorGetCursorPos(IggTextEditor handle, int* column, int* line)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  TextEditor::Coordinates col = editor->GetCursorPosition();
  *column = (float)col.mColumn;
  *line = (float)col.mLine;
}

void IggTextEditorGetSelectionStart(IggTextEditor handle, int* column, int* line)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  TextEditor::Coordinates col = editor->GetSelectionStart();
  *column = (float)col.mColumn;
  *line = (float)col.mLine;
}

void IggTextEditorSetLanguageDefinitionSQL(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetLanguageDefinition(TextEditor::LanguageDefinition::SQL());
}


void IggTextEditorSetLanguageDefinitionCPP(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetLanguageDefinition(TextEditor::LanguageDefinition::CPlusPlus());
}

void IggTextEditorSetLanguageDefinitionC(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetLanguageDefinition(TextEditor::LanguageDefinition::C());
}

void IggTextEditorSetLanguageDefinitionLua(IggTextEditor handle)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  editor->SetLanguageDefinition(TextEditor::LanguageDefinition::Lua());
}

IggTextEditorErrorMarkers IggTextEditorNewErrorMarkers() 
{
  TextEditor::ErrorMarkers *markers = new TextEditor::ErrorMarkers();
  return static_cast<IggTextEditorErrorMarkers>(markers);
}

void IggTextEditorErrorMarkersInsert(IggTextEditorErrorMarkers handle, int pos, const char* errMsg)
{
  TextEditor::ErrorMarkers *markers = reinterpret_cast<TextEditor::ErrorMarkers*>(handle);
  markers->insert(std::pair<int, std::string>(pos, errMsg));
}

void IggTextEditorErrorMarkersClear(IggTextEditorErrorMarkers marker)
{
  TextEditor::ErrorMarkers *markers = reinterpret_cast<TextEditor::ErrorMarkers*>(marker);
  markers->clear();
}


unsigned int IggTextEditorErrorMarkersSize(IggTextEditorErrorMarkers handle)
{
  TextEditor::ErrorMarkers *markers = reinterpret_cast<TextEditor::ErrorMarkers*>(handle);
  return markers->size();
}

void IggTextEditorSetErrorMarkers(IggTextEditor handle, IggTextEditorErrorMarkers marker)
{
  TextEditor *editor = reinterpret_cast<TextEditor*>(handle);
  TextEditor::ErrorMarkers *markers = reinterpret_cast<TextEditor::ErrorMarkers*>(marker);

  editor->SetErrorMarkers(*markers);
}

