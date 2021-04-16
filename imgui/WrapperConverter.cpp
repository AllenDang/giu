#include "imguiWrappedHeader.h"
#include "imguiWrapperTypes.h"
#include "WrapperConverter.h"

void importValue(bool &out, IggBool const &in)
{
   out = in != 0;
}

void exportValue(IggBool &out, bool const &in)
{
   out = in ? 1 : 0;
}

void importValue(float &out, IggFloat const &in)
{
  out = in;
}

void exportValue(IggFloat &out, float const &in)
{
  out = in;
}

void importValue(ImVec2 &out, IggVec2 const &in)
{
   out.x = in.x;
   out.y = in.y;
}

void exportValue(IggVec2 &out, ImVec2 const &in)
{
   out.x = in.x;
   out.y = in.y;
}

void importValue(ImVec4 &out, IggVec4 const &in)
{
   out.x = in.x;
   out.y = in.y;
   out.z = in.z;
   out.w = in.w;
}

void exportValue(IggVec4 &out, ImVec4 const &in)
{
   out.x = in.x;
   out.y = in.y;
   out.z = in.z;
   out.w = in.w;
}
