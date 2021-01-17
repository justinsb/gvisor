// Automatically generated marshal tests. See tools/go_marshal.

package primitive

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "gvisor.dev/gvisor/pkg/usermem"
    "gvisor.dev/gvisor/tools/go_marshal/analysis"
    "reflect"
    "testing"
)

func TestSizeNonZeroInt16(t *testing.T) {
    var x Int16
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentInt16(t *testing.T) {
    var x Int16
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataInt16(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Int16
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

func TestWriteToUnmarshalPreservesDataInt16(t *testing.T) {
    var x, y, yUnsafe Int16
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

func TestSizeBytesOnTypedNilPtrInt16(t *testing.T) {
    var x Int16
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Int16)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataInt16(t *testing.T) {
    var x, y, yUnsafe [8]Int16
    analysis.RandomizeValue(&x)

    size := (*Int16)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeInt16Slice(x[:], bufUnsafe)

    UnmarshalUnsafeInt16Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeInt16Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroInt32(t *testing.T) {
    var x Int32
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentInt32(t *testing.T) {
    var x Int32
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataInt32(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Int32
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

func TestWriteToUnmarshalPreservesDataInt32(t *testing.T) {
    var x, y, yUnsafe Int32
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

func TestSizeBytesOnTypedNilPtrInt32(t *testing.T) {
    var x Int32
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Int32)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataInt32(t *testing.T) {
    var x, y, yUnsafe [8]Int32
    analysis.RandomizeValue(&x)

    size := (*Int32)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeInt32Slice(x[:], bufUnsafe)

    UnmarshalUnsafeInt32Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeInt32Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroInt64(t *testing.T) {
    var x Int64
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentInt64(t *testing.T) {
    var x Int64
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataInt64(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Int64
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

func TestWriteToUnmarshalPreservesDataInt64(t *testing.T) {
    var x, y, yUnsafe Int64
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

func TestSizeBytesOnTypedNilPtrInt64(t *testing.T) {
    var x Int64
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Int64)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataInt64(t *testing.T) {
    var x, y, yUnsafe [8]Int64
    analysis.RandomizeValue(&x)

    size := (*Int64)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeInt64Slice(x[:], bufUnsafe)

    UnmarshalUnsafeInt64Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeInt64Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroInt8(t *testing.T) {
    var x Int8
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentInt8(t *testing.T) {
    var x Int8
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataInt8(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Int8
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

func TestWriteToUnmarshalPreservesDataInt8(t *testing.T) {
    var x, y, yUnsafe Int8
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

func TestSizeBytesOnTypedNilPtrInt8(t *testing.T) {
    var x Int8
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Int8)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataInt8(t *testing.T) {
    var x, y, yUnsafe [8]Int8
    analysis.RandomizeValue(&x)

    size := (*Int8)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeInt8Slice(x[:], bufUnsafe)

    UnmarshalUnsafeInt8Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeInt8Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroUint16(t *testing.T) {
    var x Uint16
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUint16(t *testing.T) {
    var x Uint16
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUint16(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Uint16
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

func TestWriteToUnmarshalPreservesDataUint16(t *testing.T) {
    var x, y, yUnsafe Uint16
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

func TestSizeBytesOnTypedNilPtrUint16(t *testing.T) {
    var x Uint16
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Uint16)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataUint16(t *testing.T) {
    var x, y, yUnsafe [8]Uint16
    analysis.RandomizeValue(&x)

    size := (*Uint16)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeUint16Slice(x[:], bufUnsafe)

    UnmarshalUnsafeUint16Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeUint16Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroUint32(t *testing.T) {
    var x Uint32
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUint32(t *testing.T) {
    var x Uint32
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUint32(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Uint32
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

func TestWriteToUnmarshalPreservesDataUint32(t *testing.T) {
    var x, y, yUnsafe Uint32
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

func TestSizeBytesOnTypedNilPtrUint32(t *testing.T) {
    var x Uint32
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Uint32)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataUint32(t *testing.T) {
    var x, y, yUnsafe [8]Uint32
    analysis.RandomizeValue(&x)

    size := (*Uint32)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeUint32Slice(x[:], bufUnsafe)

    UnmarshalUnsafeUint32Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeUint32Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroUint64(t *testing.T) {
    var x Uint64
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUint64(t *testing.T) {
    var x Uint64
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUint64(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Uint64
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

func TestWriteToUnmarshalPreservesDataUint64(t *testing.T) {
    var x, y, yUnsafe Uint64
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

func TestSizeBytesOnTypedNilPtrUint64(t *testing.T) {
    var x Uint64
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Uint64)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataUint64(t *testing.T) {
    var x, y, yUnsafe [8]Uint64
    analysis.RandomizeValue(&x)

    size := (*Uint64)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeUint64Slice(x[:], bufUnsafe)

    UnmarshalUnsafeUint64Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeUint64Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroUint8(t *testing.T) {
    var x Uint8
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUint8(t *testing.T) {
    var x Uint8
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUint8(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Uint8
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

func TestWriteToUnmarshalPreservesDataUint8(t *testing.T) {
    var x, y, yUnsafe Uint8
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

func TestSizeBytesOnTypedNilPtrUint8(t *testing.T) {
    var x Uint8
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Uint8)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataUint8(t *testing.T) {
    var x, y, yUnsafe [8]Uint8
    analysis.RandomizeValue(&x)

    size := (*Uint8)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeUint8Slice(x[:], bufUnsafe)

    UnmarshalUnsafeUint8Slice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeUint8Slice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

