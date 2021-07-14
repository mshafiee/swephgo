// Copyright 2021 Mohammad Shafiee and The SwEphGo Authors
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/gpl-3.0.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swephgo

/*
#include "swephexp.h"
#include "sweph.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Int32 type as declared in src/sweodef.h:210
type Int32 int32

// Int64 type as declared in src/sweodef.h:211
type Int64 int64

// Uint32 type as declared in src/sweodef.h:212
type Uint32 uint32

// Int16 type as declared in src/sweodef.h:213
type Int16 int16

// REAL8 type as declared in src/sweodef.h:214
type REAL8 float64

// INT4 type as declared in src/sweodef.h:215
type INT4 int32

// UINT4 type as declared in src/sweodef.h:216
type UINT4 uint32

// ASBOOL type as declared in src/sweodef.h:218
type ASBOOL int32

// UINT2 type as declared in src/sweodef.h:219
type UINT2 uint16

// UCHAR type as declared in src/sweodef.h:249
type UCHAR byte

// Centisec type as declared in src/sweodef.h:275
type Centisec int32

// SweData as declared in src/sweph.h:794
type SweData struct {
	EphePathIsSet      int32
	JplFileIsOpen      int32
	Fixfp              []FILE
	Ephepath           [256]byte
	Jplfnam            [256]byte
	Jpldenum           int32
	LastEpheflag       int32
	GeoposIsSet        int32
	AyanaIsSet         int32
	IsOldStarfile      int32
	EopTjdBeg          float64
	EopTjdBegHorizons  float64
	EopTjdEnd          float64
	EopTjdEndAdd       float64
	EopDpsiLoaded      int32
	TidAcc             float64
	IsTidAccManual     int32
	InitDtDone         int32
	SwedIsInitialised  int32
	DeltaTUserdefIsSet int32
	DeltaTUserdef      float64
	AstG               float64
	AstH               float64
	AstDiam            float64
	Astelem            [2560]byte
	ISavedPlanetName   int32
	SavedPlanetName    [80]byte
	Dpsi               []float64
	Deps               []float64
	Timeout            int32
	AstroModels        [8]int32
	DoInterpolateNut   int32
	Interpol           Interpol
	Fidat              [7]FileData
	Gcdat              GenConst
	Pldat              [18]PlanData
	Nddat              [6]PlanData
	Savedat            [24]SavePositions
	Oec                Epsilon
	Oec2000            Epsilon
	Nut                Nut
	Nut2000            Nut
	Nutv               Nut
	Topd               TopoData
	Sidd               SidData
	NFixstarsReal      int32
	NFixstarsNamed     int32
	NFixstarsRecords   int32
	FixedStars         []FixedStar
	ref84c87565        *C.struct_swe_data
	allocs84c87565     interface{}
}

// Interpol as declared in src/sweph.h:787
type Interpol struct {
	TjdNut0       float64
	TjdNut2       float64
	NutDpsi0      float64
	NutDpsi1      float64
	NutDpsi2      float64
	NutDeps0      float64
	NutDeps1      float64
	NutDeps2      float64
	ref7ff86a1    *C.struct_interpol
	allocs7ff86a1 interface{}
}

// FileData as declared in src/sweph.h:711
type FileData struct {
	Fnam           [256]byte
	Fversion       int32
	Astnam         [50]byte
	SwephDenum     int32
	Fptr           []FILE
	Tfstart        float64
	Tfend          float64
	Iflg           int32
	Npl            int16
	Ipl            [50]int32
	refa68358f4    *C.struct_file_data
	allocsa68358f4 interface{}
}

// GenConst as declared in src/sweph.h:725
type GenConst struct {
	Clight         float64
	Aunit          float64
	Helgravconst   float64
	Ratme          float64
	Sunradius      float64
	ref4e398b8d    *C.struct_gen_const
	allocs4e398b8d interface{}
}

// PlanData as declared in src/sweph.h:613
type PlanData struct {
	Ibdy           int32
	Iflg           int32
	Ncoe           int32
	Lndx0          int32
	Nndx           int32
	Tfstart        float64
	Tfend          float64
	Dseg           float64
	Telem          float64
	Prot           float64
	Qrot           float64
	Dprot          float64
	Dqrot          float64
	Rmax           float64
	Peri           float64
	Dperi          float64
	Refep          []float64
	Tseg0          float64
	Tseg1          float64
	Segp           []float64
	Neval          int32
	Teval          float64
	Iephe          int32
	X              [6]float64
	Xflgs          int32
	Xreturn        [24]float64
	ref3aa9fc68    *C.struct_plan_data
	allocs3aa9fc68 interface{}
}

// SavePositions as declared in src/sweph.h:733
type SavePositions struct {
	Ipl            int32
	Tsave          float64
	Iflgsave       int32
	Xsaves         [24]float64
	refc6ebb512    *C.struct_save_positions
	allocsc6ebb512 interface{}
}

// Epsilon as declared in src/sweph.h:604
type Epsilon struct {
	Teps           float64
	Eps            float64
	Seps           float64
	Ceps           float64
	ref3747120b    *C.struct_epsilon
	allocs3747120b interface{}
}

// Nut as declared in src/sweph.h:694
type Nut struct {
	Tnut           float64
	Nutlo          [2]float64
	Snut           float64
	Cnut           float64
	Matrix         [3][3]float64
	ref66f7e72c    *C.struct_nut
	allocs66f7e72c interface{}
}

// TopoData as declared in src/sweph.h:761
type TopoData struct {
	Geolon         float64
	Geolat         float64
	Geoalt         float64
	Teval          float64
	TjdUt          float64
	Xobs           [6]float64
	refe33fee58    *C.struct_topo_data
	allocse33fee58 interface{}
}

// SidData as declared in src/sweph.h:768
type SidData struct {
	SidMode        int32
	AyanT0         float64
	T0             float64
	T0IsUT         int32
	ref78754153    *C.struct_sid_data
	allocs78754153 interface{}
}

// FixedStar as declared in src/sweph.h:776
type FixedStar struct {
	Skey           [42]byte
	Starname       [41]byte
	Starbayer      [41]byte
	Starno         [10]byte
	Epoch          float64
	Ra             float64
	De             float64
	Ramot          float64
	Demot          float64
	Radvel         float64
	Parall         float64
	Mag            float64
	ref6cd294fc    *C.struct_fixed_star
	allocs6cd294fc interface{}
}

// FILE as declared in include/_stdio.h:157
type FILE struct {
	P              []byte
	R              int32
	W              int32
	Flags          int16
	File           int16
	Bf             _Sbuf
	Lbfsize        int32
	Cookie         unsafe.Pointer
	Close          *func(arg0 unsafe.Pointer) int32
	Read           *func(arg0 unsafe.Pointer, arg1 []byte, arg2 int32) int32
	Seek           FposT
	Write          *func(arg0 unsafe.Pointer, arg1 []string, arg2 int32) int32
	Ub             _Sbuf
	Extra          []_SFILEX
	Ur             int32
	Ubuf           [3]byte
	Nbuf           [1]byte
	Lb             _Sbuf
	Blksize        int32
	Offset         FposT
	refba0adba4    *C.FILE
	allocsba0adba4 interface{}
}

// _Sbuf as declared in include/_stdio.h:92
type _Sbuf struct {
	Base           []byte
	Size           int32
	ref98bae9e0    *C.struct___sbuf
	allocs98bae9e0 interface{}
}

// FposT type as declared in include/_stdio.h:81
type FposT int64

// _SFILEX as declared in include/_stdio.h:98
type _SFILEX C.struct___sFILEX
