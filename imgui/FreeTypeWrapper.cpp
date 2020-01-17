#ifdef IMGUI_FREETYPE_ENABLED

#include "imguiWrappedHeader.h"
#include "imgui_freetype.h"
#include "FreeTypeWrapper.h"

int iggFreeTypeBuildFontAtlas(IggFontAtlas handle, unsigned int flags)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return ImGuiFreeType::BuildFontAtlas(fontAtlas, flags);
}

#endif
