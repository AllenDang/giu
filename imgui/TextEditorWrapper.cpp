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
