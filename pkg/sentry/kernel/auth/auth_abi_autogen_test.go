// Automatically generated marshal tests. See tools/go_marshal.

package auth

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "gvisor.dev/gvisor/pkg/usermem"
    "gvisor.dev/gvisor/tools/go_marshal/analysis"
    "reflect"
    "testing"
)

func TestSizeNonZeroGID(t *testing.T) {
    var x GID
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentGID(t *testing.T) {
    var x GID
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataGID(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe GID
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

func TestWriteToUnmarshalPreservesDataGID(t *testing.T) {
    var x, y, yUnsafe GID
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

func TestSizeBytesOnTypedNilPtrGID(t *testing.T) {
    var x GID
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*GID)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataGID(t *testing.T) {
    var x, y, yUnsafe [8]GID
    analysis.RandomizeValue(&x)

    size := (*GID)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeGIDSlice(x[:], bufUnsafe)

    UnmarshalUnsafeGIDSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeGIDSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroUID(t *testing.T) {
    var x UID
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUID(t *testing.T) {
    var x UID
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUID(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe UID
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

func TestWriteToUnmarshalPreservesDataUID(t *testing.T) {
    var x, y, yUnsafe UID
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

func TestSizeBytesOnTypedNilPtrUID(t *testing.T) {
    var x UID
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*UID)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

