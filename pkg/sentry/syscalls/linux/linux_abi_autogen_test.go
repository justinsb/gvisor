// Automatically generated marshal tests. See tools/go_marshal.

package linux

import (
    "bytes"
    "fmt"
    "gvisor.dev/gvisor/tools/go_marshal/analysis"
    "reflect"
    "testing"
)

func TestSizeNonZeroDirentHdr(t *testing.T) {
    var x direntHdr
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentDirentHdr(t *testing.T) {
    var x direntHdr
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataDirentHdr(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe direntHdr
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

func TestWriteToUnmarshalPreservesDataDirentHdr(t *testing.T) {
    var x, y, yUnsafe direntHdr
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

func TestSizeBytesOnTypedNilPtrDirentHdr(t *testing.T) {
    var x direntHdr
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*direntHdr)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroOldDirentHdr(t *testing.T) {
    var x oldDirentHdr
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentOldDirentHdr(t *testing.T) {
    var x oldDirentHdr
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataOldDirentHdr(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe oldDirentHdr
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

func TestWriteToUnmarshalPreservesDataOldDirentHdr(t *testing.T) {
    var x, y, yUnsafe oldDirentHdr
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

func TestSizeBytesOnTypedNilPtrOldDirentHdr(t *testing.T) {
    var x oldDirentHdr
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*oldDirentHdr)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroRlimit64(t *testing.T) {
    var x rlimit64
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentRlimit64(t *testing.T) {
    var x rlimit64
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataRlimit64(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe rlimit64
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

func TestWriteToUnmarshalPreservesDataRlimit64(t *testing.T) {
    var x, y, yUnsafe rlimit64
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

func TestSizeBytesOnTypedNilPtrRlimit64(t *testing.T) {
    var x rlimit64
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*rlimit64)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSchedParam(t *testing.T) {
    var x SchedParam
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSchedParam(t *testing.T) {
    var x SchedParam
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSchedParam(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SchedParam
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

func TestWriteToUnmarshalPreservesDataSchedParam(t *testing.T) {
    var x, y, yUnsafe SchedParam
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

func TestSizeBytesOnTypedNilPtrSchedParam(t *testing.T) {
    var x SchedParam
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SchedParam)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroUserSockFprog(t *testing.T) {
    var x userSockFprog
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUserSockFprog(t *testing.T) {
    var x userSockFprog
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUserSockFprog(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe userSockFprog
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

func TestWriteToUnmarshalPreservesDataUserSockFprog(t *testing.T) {
    var x, y, yUnsafe userSockFprog
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

func TestSizeBytesOnTypedNilPtrUserSockFprog(t *testing.T) {
    var x userSockFprog
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*userSockFprog)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroMessageHeader64(t *testing.T) {
    var x MessageHeader64
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentMessageHeader64(t *testing.T) {
    var x MessageHeader64
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataMessageHeader64(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe MessageHeader64
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

func TestWriteToUnmarshalPreservesDataMessageHeader64(t *testing.T) {
    var x, y, yUnsafe MessageHeader64
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

func TestSizeBytesOnTypedNilPtrMessageHeader64(t *testing.T) {
    var x MessageHeader64
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*MessageHeader64)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroMultipleMessageHeader64(t *testing.T) {
    var x multipleMessageHeader64
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentMultipleMessageHeader64(t *testing.T) {
    var x multipleMessageHeader64
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataMultipleMessageHeader64(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe multipleMessageHeader64
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

func TestWriteToUnmarshalPreservesDataMultipleMessageHeader64(t *testing.T) {
    var x, y, yUnsafe multipleMessageHeader64
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

func TestSizeBytesOnTypedNilPtrMultipleMessageHeader64(t *testing.T) {
    var x multipleMessageHeader64
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*multipleMessageHeader64)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

