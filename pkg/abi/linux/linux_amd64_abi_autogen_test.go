// Automatically generated marshal tests. See tools/go_marshal.

// +build amd64
// +build amd64
// +build amd64
// +build amd64
// +build amd64

package linux

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "gvisor.dev/gvisor/pkg/usermem"
    "gvisor.dev/gvisor/tools/go_marshal/analysis"
    "reflect"
    "testing"
)

func TestSizeNonZeroEpollEvent(t *testing.T) {
    var x EpollEvent
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentEpollEvent(t *testing.T) {
    var x EpollEvent
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataEpollEvent(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe EpollEvent
    analysis.RandomizeValue(&x)

    buf := make([]byte, x.SizeBytes())
    x.MarshalBytes(buf)
    bufUnsafe := make([]byte, x.SizeBytes())
    x.MarshalUnsafe(bufUnsafe)

    y.UnmarshalBytes(buf)
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    yUnsafe.UnmarshalBytes(bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

    z.UnmarshalUnsafe(buf)
    if !reflect.DeepEqual(x, z) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, z))
    }
    zUnsafe.UnmarshalUnsafe(bufUnsafe)
    if !reflect.DeepEqual(x, zUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, zUnsafe))
    }
}

func TestWriteToUnmarshalPreservesDataEpollEvent(t *testing.T) {
    var x, y, yUnsafe EpollEvent
    analysis.RandomizeValue(&x)

    var buf bytes.Buffer

    x.WriteTo(&buf)
    y.UnmarshalBytes(buf.Bytes())

    yUnsafe.UnmarshalUnsafe(buf.Bytes())

    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }
}

func TestSizeBytesOnTypedNilPtrEpollEvent(t *testing.T) {
    var x EpollEvent
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*EpollEvent)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataEpollEvent(t *testing.T) {
    var x, y, yUnsafe [8]EpollEvent
    analysis.RandomizeValue(&x)

    size := (*EpollEvent)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeEpollEventSlice(x[:], bufUnsafe)

    UnmarshalUnsafeEpollEventSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeEpollEventSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroStat(t *testing.T) {
    var x Stat
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentStat(t *testing.T) {
    var x Stat
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataStat(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Stat
    analysis.RandomizeValue(&x)

    buf := make([]byte, x.SizeBytes())
    x.MarshalBytes(buf)
    bufUnsafe := make([]byte, x.SizeBytes())
    x.MarshalUnsafe(bufUnsafe)

    y.UnmarshalBytes(buf)
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    yUnsafe.UnmarshalBytes(bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

    z.UnmarshalUnsafe(buf)
    if !reflect.DeepEqual(x, z) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, z))
    }
    zUnsafe.UnmarshalUnsafe(bufUnsafe)
    if !reflect.DeepEqual(x, zUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, zUnsafe))
    }
}

func TestWriteToUnmarshalPreservesDataStat(t *testing.T) {
    var x, y, yUnsafe Stat
    analysis.RandomizeValue(&x)

    var buf bytes.Buffer

    x.WriteTo(&buf)
    y.UnmarshalBytes(buf.Bytes())

    yUnsafe.UnmarshalUnsafe(buf.Bytes())

    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }
}

func TestSizeBytesOnTypedNilPtrStat(t *testing.T) {
    var x Stat
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Stat)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroPtraceRegs(t *testing.T) {
    var x PtraceRegs
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentPtraceRegs(t *testing.T) {
    var x PtraceRegs
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataPtraceRegs(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe PtraceRegs
    analysis.RandomizeValue(&x)

    buf := make([]byte, x.SizeBytes())
    x.MarshalBytes(buf)
    bufUnsafe := make([]byte, x.SizeBytes())
    x.MarshalUnsafe(bufUnsafe)

    y.UnmarshalBytes(buf)
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    yUnsafe.UnmarshalBytes(bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

    z.UnmarshalUnsafe(buf)
    if !reflect.DeepEqual(x, z) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, z))
    }
    zUnsafe.UnmarshalUnsafe(bufUnsafe)
    if !reflect.DeepEqual(x, zUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, zUnsafe))
    }
}

func TestWriteToUnmarshalPreservesDataPtraceRegs(t *testing.T) {
    var x, y, yUnsafe PtraceRegs
    analysis.RandomizeValue(&x)

    var buf bytes.Buffer

    x.WriteTo(&buf)
    y.UnmarshalBytes(buf.Bytes())

    yUnsafe.UnmarshalUnsafe(buf.Bytes())

    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }
}

func TestSizeBytesOnTypedNilPtrPtraceRegs(t *testing.T) {
    var x PtraceRegs
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*PtraceRegs)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSemidDS(t *testing.T) {
    var x SemidDS
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSemidDS(t *testing.T) {
    var x SemidDS
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSemidDS(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SemidDS
    analysis.RandomizeValue(&x)

    buf := make([]byte, x.SizeBytes())
    x.MarshalBytes(buf)
    bufUnsafe := make([]byte, x.SizeBytes())
    x.MarshalUnsafe(bufUnsafe)

    y.UnmarshalBytes(buf)
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    yUnsafe.UnmarshalBytes(bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

    z.UnmarshalUnsafe(buf)
    if !reflect.DeepEqual(x, z) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalBytes/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, z))
    }
    zUnsafe.UnmarshalUnsafe(bufUnsafe)
    if !reflect.DeepEqual(x, zUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafe/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, zUnsafe))
    }
}

func TestWriteToUnmarshalPreservesDataSemidDS(t *testing.T) {
    var x, y, yUnsafe SemidDS
    analysis.RandomizeValue(&x)

    var buf bytes.Buffer

    x.WriteTo(&buf)
    y.UnmarshalBytes(buf.Bytes())

    yUnsafe.UnmarshalUnsafe(buf.Bytes())

    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalBytes cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across WriteTo/UnmarshalUnsafe cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }
}

func TestSizeBytesOnTypedNilPtrSemidDS(t *testing.T) {
    var x SemidDS
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SemidDS)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

