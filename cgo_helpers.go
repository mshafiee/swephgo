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
#cgo CFLAGS: -I.
#cgo LDFLAGS: -lswe
#include "swephexp.h"
#include "sweph.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocStructSweDataMemory allocates memory for type C.struct_swe_data in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructSweDataMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructSweDataValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfStructSweDataValue = unsafe.Sizeof([1]C.struct_swe_data{})

// copyPDoubleBytes copies the data from Go slice as *C.double.
func copyPDoubleBytes(slice *sliceHeader) (*C.double, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	//mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
	//	Data: slice.Data,
	//	Len:  int(sizeOfDoubleValue) * slice.Len,
	//	Cap:  int(sizeOfDoubleValue) * slice.Len,
	//}))))
	mem0 := unsafe.Pointer((*sliceHeader)(unsafe.Pointer(slice)).Data)
	//allocs.Add(mem0)

	return (*C.double)(mem0), allocs
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// allocDoubleMemory allocates memory for type C.double in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDoubleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDoubleValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfDoubleValue = unsafe.Sizeof([1]C.double{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SweData) Ref() *C.struct_swe_data {
	if x == nil {
		return nil
	}
	return x.ref84c87565
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SweData) Free() {
	if x != nil && x.allocs84c87565 != nil {
		x.allocs84c87565.(*cgoAllocMap).Free()
		x.ref84c87565 = nil
	}
}

// NewSweDataRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSweDataRef(ref unsafe.Pointer) *SweData {
	if ref == nil {
		return nil
	}
	obj := new(SweData)
	obj.ref84c87565 = (*C.struct_swe_data)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SweData) PassRef() (*C.struct_swe_data, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref84c87565 != nil {
		return x.ref84c87565, nil
	}
	mem84c87565 := allocStructSweDataMemory(1)
	ref84c87565 := (*C.struct_swe_data)(mem84c87565)
	allocs84c87565 := new(cgoAllocMap)
	allocs84c87565.Add(mem84c87565)

	var cephe_path_is_set_allocs *cgoAllocMap
	ref84c87565.ephe_path_is_set, cephe_path_is_set_allocs = (C.AS_BOOL)(x.EphePathIsSet), cgoAllocsUnknown
	allocs84c87565.Borrow(cephe_path_is_set_allocs)

	var cjpl_file_is_open_allocs *cgoAllocMap
	ref84c87565.jpl_file_is_open, cjpl_file_is_open_allocs = (C.AS_BOOL)(x.JplFileIsOpen), cgoAllocsUnknown
	allocs84c87565.Borrow(cjpl_file_is_open_allocs)

	var cephepath_allocs *cgoAllocMap
	ref84c87565.ephepath, cephepath_allocs = *(*[256]C.char)(unsafe.Pointer(&x.Ephepath)), cgoAllocsUnknown
	allocs84c87565.Borrow(cephepath_allocs)

	var cjplfnam_allocs *cgoAllocMap
	ref84c87565.jplfnam, cjplfnam_allocs = *(*[256]C.char)(unsafe.Pointer(&x.Jplfnam)), cgoAllocsUnknown
	allocs84c87565.Borrow(cjplfnam_allocs)

	var cjpldenum_allocs *cgoAllocMap
	ref84c87565.jpldenum, cjpldenum_allocs = (C.int32)(x.Jpldenum), cgoAllocsUnknown
	allocs84c87565.Borrow(cjpldenum_allocs)

	var clast_epheflag_allocs *cgoAllocMap
	ref84c87565.last_epheflag, clast_epheflag_allocs = (C.int32)(x.LastEpheflag), cgoAllocsUnknown
	allocs84c87565.Borrow(clast_epheflag_allocs)

	var cgeopos_is_set_allocs *cgoAllocMap
	ref84c87565.geopos_is_set, cgeopos_is_set_allocs = (C.AS_BOOL)(x.GeoposIsSet), cgoAllocsUnknown
	allocs84c87565.Borrow(cgeopos_is_set_allocs)

	var cayana_is_set_allocs *cgoAllocMap
	ref84c87565.ayana_is_set, cayana_is_set_allocs = (C.AS_BOOL)(x.AyanaIsSet), cgoAllocsUnknown
	allocs84c87565.Borrow(cayana_is_set_allocs)

	var cis_old_starfile_allocs *cgoAllocMap
	ref84c87565.is_old_starfile, cis_old_starfile_allocs = (C.AS_BOOL)(x.IsOldStarfile), cgoAllocsUnknown
	allocs84c87565.Borrow(cis_old_starfile_allocs)

	var ceop_tjd_beg_allocs *cgoAllocMap
	ref84c87565.eop_tjd_beg, ceop_tjd_beg_allocs = (C.double)(x.EopTjdBeg), cgoAllocsUnknown
	allocs84c87565.Borrow(ceop_tjd_beg_allocs)

	var ceop_tjd_beg_horizons_allocs *cgoAllocMap
	ref84c87565.eop_tjd_beg_horizons, ceop_tjd_beg_horizons_allocs = (C.double)(x.EopTjdBegHorizons), cgoAllocsUnknown
	allocs84c87565.Borrow(ceop_tjd_beg_horizons_allocs)

	var ceop_tjd_end_allocs *cgoAllocMap
	ref84c87565.eop_tjd_end, ceop_tjd_end_allocs = (C.double)(x.EopTjdEnd), cgoAllocsUnknown
	allocs84c87565.Borrow(ceop_tjd_end_allocs)

	var ceop_tjd_end_add_allocs *cgoAllocMap
	ref84c87565.eop_tjd_end_add, ceop_tjd_end_add_allocs = (C.double)(x.EopTjdEndAdd), cgoAllocsUnknown
	allocs84c87565.Borrow(ceop_tjd_end_add_allocs)

	var ceop_dpsi_loaded_allocs *cgoAllocMap
	ref84c87565.eop_dpsi_loaded, ceop_dpsi_loaded_allocs = (C.int)(x.EopDpsiLoaded), cgoAllocsUnknown
	allocs84c87565.Borrow(ceop_dpsi_loaded_allocs)

	var ctid_acc_allocs *cgoAllocMap
	ref84c87565.tid_acc, ctid_acc_allocs = (C.double)(x.TidAcc), cgoAllocsUnknown
	allocs84c87565.Borrow(ctid_acc_allocs)

	var cis_tid_acc_manual_allocs *cgoAllocMap
	ref84c87565.is_tid_acc_manual, cis_tid_acc_manual_allocs = (C.AS_BOOL)(x.IsTidAccManual), cgoAllocsUnknown
	allocs84c87565.Borrow(cis_tid_acc_manual_allocs)

	var cinit_dt_done_allocs *cgoAllocMap
	ref84c87565.init_dt_done, cinit_dt_done_allocs = (C.AS_BOOL)(x.InitDtDone), cgoAllocsUnknown
	allocs84c87565.Borrow(cinit_dt_done_allocs)

	var cswed_is_initialised_allocs *cgoAllocMap
	ref84c87565.swed_is_initialised, cswed_is_initialised_allocs = (C.AS_BOOL)(x.SwedIsInitialised), cgoAllocsUnknown
	allocs84c87565.Borrow(cswed_is_initialised_allocs)

	var cdelta_t_userdef_is_set_allocs *cgoAllocMap
	ref84c87565.delta_t_userdef_is_set, cdelta_t_userdef_is_set_allocs = (C.AS_BOOL)(x.DeltaTUserdefIsSet), cgoAllocsUnknown
	allocs84c87565.Borrow(cdelta_t_userdef_is_set_allocs)

	var cdelta_t_userdef_allocs *cgoAllocMap
	ref84c87565.delta_t_userdef, cdelta_t_userdef_allocs = (C.double)(x.DeltaTUserdef), cgoAllocsUnknown
	allocs84c87565.Borrow(cdelta_t_userdef_allocs)

	var cast_G_allocs *cgoAllocMap
	ref84c87565.ast_G, cast_G_allocs = (C.double)(x.AstG), cgoAllocsUnknown
	allocs84c87565.Borrow(cast_G_allocs)

	var cast_H_allocs *cgoAllocMap
	ref84c87565.ast_H, cast_H_allocs = (C.double)(x.AstH), cgoAllocsUnknown
	allocs84c87565.Borrow(cast_H_allocs)

	var cast_diam_allocs *cgoAllocMap
	ref84c87565.ast_diam, cast_diam_allocs = (C.double)(x.AstDiam), cgoAllocsUnknown
	allocs84c87565.Borrow(cast_diam_allocs)

	var castelem_allocs *cgoAllocMap
	ref84c87565.astelem, castelem_allocs = *(*[2560]C.char)(unsafe.Pointer(&x.Astelem)), cgoAllocsUnknown
	allocs84c87565.Borrow(castelem_allocs)

	var ci_saved_planet_name_allocs *cgoAllocMap
	ref84c87565.i_saved_planet_name, ci_saved_planet_name_allocs = (C.int)(x.ISavedPlanetName), cgoAllocsUnknown
	allocs84c87565.Borrow(ci_saved_planet_name_allocs)

	var csaved_planet_name_allocs *cgoAllocMap
	ref84c87565.saved_planet_name, csaved_planet_name_allocs = *(*[80]C.char)(unsafe.Pointer(&x.SavedPlanetName)), cgoAllocsUnknown
	allocs84c87565.Borrow(csaved_planet_name_allocs)

	var cdpsi_allocs *cgoAllocMap
	ref84c87565.dpsi, cdpsi_allocs = copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&x.Dpsi)))
	allocs84c87565.Borrow(cdpsi_allocs)

	var cdeps_allocs *cgoAllocMap
	ref84c87565.deps, cdeps_allocs = copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&x.Deps)))
	allocs84c87565.Borrow(cdeps_allocs)

	var ctimeout_allocs *cgoAllocMap
	ref84c87565.timeout, ctimeout_allocs = (C.int32)(x.Timeout), cgoAllocsUnknown
	allocs84c87565.Borrow(ctimeout_allocs)

	var castro_models_allocs *cgoAllocMap
	ref84c87565.astro_models, castro_models_allocs = *(*[8]C.int32)(unsafe.Pointer(&x.AstroModels)), cgoAllocsUnknown
	allocs84c87565.Borrow(castro_models_allocs)

	var cdo_interpolate_nut_allocs *cgoAllocMap
	ref84c87565.do_interpolate_nut, cdo_interpolate_nut_allocs = (C.AS_BOOL)(x.DoInterpolateNut), cgoAllocsUnknown
	allocs84c87565.Borrow(cdo_interpolate_nut_allocs)

	var cn_fixstars_real_allocs *cgoAllocMap
	ref84c87565.n_fixstars_real, cn_fixstars_real_allocs = (C.AS_BOOL)(x.NFixstarsReal), cgoAllocsUnknown
	allocs84c87565.Borrow(cn_fixstars_real_allocs)

	var cn_fixstars_named_allocs *cgoAllocMap
	ref84c87565.n_fixstars_named, cn_fixstars_named_allocs = (C.AS_BOOL)(x.NFixstarsNamed), cgoAllocsUnknown
	allocs84c87565.Borrow(cn_fixstars_named_allocs)

	var cn_fixstars_records_allocs *cgoAllocMap
	ref84c87565.n_fixstars_records, cn_fixstars_records_allocs = (C.AS_BOOL)(x.NFixstarsRecords), cgoAllocsUnknown
	allocs84c87565.Borrow(cn_fixstars_records_allocs)

	x.ref84c87565 = ref84c87565
	x.allocs84c87565 = allocs84c87565
	return ref84c87565, allocs84c87565

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SweData) PassValue() (C.struct_swe_data, *cgoAllocMap) {
	if x.ref84c87565 != nil {
		return *x.ref84c87565, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SweData) Deref() {
	if x.ref84c87565 == nil {
		return
	}
	x.EphePathIsSet = (int32)(x.ref84c87565.ephe_path_is_set)
	x.JplFileIsOpen = (int32)(x.ref84c87565.jpl_file_is_open)
	x.Ephepath = *(*[256]byte)(unsafe.Pointer(&x.ref84c87565.ephepath))
	x.Jplfnam = *(*[256]byte)(unsafe.Pointer(&x.ref84c87565.jplfnam))
	x.Jpldenum = (int32)(x.ref84c87565.jpldenum)
	x.LastEpheflag = (int32)(x.ref84c87565.last_epheflag)
	x.GeoposIsSet = (int32)(x.ref84c87565.geopos_is_set)
	x.AyanaIsSet = (int32)(x.ref84c87565.ayana_is_set)
	x.IsOldStarfile = (int32)(x.ref84c87565.is_old_starfile)
	x.EopTjdBeg = (float64)(x.ref84c87565.eop_tjd_beg)
	x.EopTjdBegHorizons = (float64)(x.ref84c87565.eop_tjd_beg_horizons)
	x.EopTjdEnd = (float64)(x.ref84c87565.eop_tjd_end)
	x.EopTjdEndAdd = (float64)(x.ref84c87565.eop_tjd_end_add)
	x.EopDpsiLoaded = (int32)(x.ref84c87565.eop_dpsi_loaded)
	x.TidAcc = (float64)(x.ref84c87565.tid_acc)
	x.IsTidAccManual = (int32)(x.ref84c87565.is_tid_acc_manual)
	x.InitDtDone = (int32)(x.ref84c87565.init_dt_done)
	x.SwedIsInitialised = (int32)(x.ref84c87565.swed_is_initialised)
	x.DeltaTUserdefIsSet = (int32)(x.ref84c87565.delta_t_userdef_is_set)
	x.DeltaTUserdef = (float64)(x.ref84c87565.delta_t_userdef)
	x.AstG = (float64)(x.ref84c87565.ast_G)
	x.AstH = (float64)(x.ref84c87565.ast_H)
	x.AstDiam = (float64)(x.ref84c87565.ast_diam)
	x.Astelem = *(*[2560]byte)(unsafe.Pointer(&x.ref84c87565.astelem))
	x.ISavedPlanetName = (int32)(x.ref84c87565.i_saved_planet_name)
	x.SavedPlanetName = *(*[80]byte)(unsafe.Pointer(&x.ref84c87565.saved_planet_name))
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.Dpsi))
	hxfc4425b.Data = unsafe.Pointer(x.ref84c87565.dpsi)
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.Deps))
	hxf95e7c8.Data = unsafe.Pointer(x.ref84c87565.deps)
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	x.Timeout = (int32)(x.ref84c87565.timeout)
	x.AstroModels = *(*[8]int32)(unsafe.Pointer(&x.ref84c87565.astro_models))
	x.DoInterpolateNut = (int32)(x.ref84c87565.do_interpolate_nut)
	x.NFixstarsReal = (int32)(x.ref84c87565.n_fixstars_real)
	x.NFixstarsNamed = (int32)(x.ref84c87565.n_fixstars_named)
	x.NFixstarsRecords = (int32)(x.ref84c87565.n_fixstars_records)
}

// copyPCharBytes copies the data from Go slice as *C.char.
func copyPCharBytes(slice *sliceHeader) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	//mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
	//	Data: slice.Data,
	//	Len:  int(sizeOfCharValue) * slice.Len,
	//	Cap:  int(sizeOfCharValue) * slice.Len,
	//}))))
	mem0 := unsafe.Pointer((*sliceHeader)(unsafe.Pointer(slice)).Data)
	//allocs.Add(mem0)

	return (*C.char)(mem0), allocs
}

// allocCharMemory allocates memory for type C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfCharValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfCharValue = unsafe.Sizeof([1]C.char{})

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

// copyPIntBytes copies the data from Go slice as *C.int.
func copyPIntBytes(slice *sliceHeader) (*C.int, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	//mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
	//	Data: slice.Data,
	//	Len:  int(sizeOfIntValue) * slice.Len,
	//	Cap:  int(sizeOfIntValue) * slice.Len,
	//}))))
	mem0 := unsafe.Pointer((*sliceHeader)(unsafe.Pointer(slice)).Data)
	//allocs.Add(mem0)

	return (*C.int)(mem0), allocs
}

// allocIntMemory allocates memory for type C.int in C.
// The caller is responsible for freeing the this memory via C.free.
func allocIntMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfIntValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfIntValue = unsafe.Sizeof([1]C.int{})

// copyPInt32Bytes copies the data from Go slice as *C.int32.
func copyPInt32Bytes(slice *sliceHeader) (*C.int32, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	//mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
	//	Data: slice.Data,
	//	Len:  int(sizeOfInt32Value) * slice.Len,
	//	Cap:  int(sizeOfInt32Value) * slice.Len,
	//}))))
	mem0 := unsafe.Pointer((*sliceHeader)(unsafe.Pointer(slice)).Data)
	//allocs.Add(mem0)

	return (*C.int32)(mem0), allocs
}

// allocInt32Memory allocates memory for type C.int32 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocInt32Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfInt32Value))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfInt32Value = unsafe.Sizeof([1]C.int32{})
