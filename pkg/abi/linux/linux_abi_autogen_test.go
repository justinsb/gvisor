// Automatically generated marshal tests. See tools/go_marshal.

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

func TestSizeNonZeroIOCallback(t *testing.T) {
    var x IOCallback
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIOCallback(t *testing.T) {
    var x IOCallback
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIOCallback(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IOCallback
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

func TestWriteToUnmarshalPreservesDataIOCallback(t *testing.T) {
    var x, y, yUnsafe IOCallback
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

func TestSizeBytesOnTypedNilPtrIOCallback(t *testing.T) {
    var x IOCallback
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IOCallback)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIOEvent(t *testing.T) {
    var x IOEvent
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIOEvent(t *testing.T) {
    var x IOEvent
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIOEvent(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IOEvent
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

func TestWriteToUnmarshalPreservesDataIOEvent(t *testing.T) {
    var x, y, yUnsafe IOEvent
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

func TestSizeBytesOnTypedNilPtrIOEvent(t *testing.T) {
    var x IOEvent
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IOEvent)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroBPFInstruction(t *testing.T) {
    var x BPFInstruction
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentBPFInstruction(t *testing.T) {
    var x BPFInstruction
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataBPFInstruction(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe BPFInstruction
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

func TestWriteToUnmarshalPreservesDataBPFInstruction(t *testing.T) {
    var x, y, yUnsafe BPFInstruction
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

func TestSizeBytesOnTypedNilPtrBPFInstruction(t *testing.T) {
    var x BPFInstruction
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*BPFInstruction)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataBPFInstruction(t *testing.T) {
    var x, y, yUnsafe [8]BPFInstruction
    analysis.RandomizeValue(&x)

    size := (*BPFInstruction)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeBPFInstructionSlice(x[:], bufUnsafe)

    UnmarshalUnsafeBPFInstructionSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeBPFInstructionSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroCapUserData(t *testing.T) {
    var x CapUserData
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentCapUserData(t *testing.T) {
    var x CapUserData
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataCapUserData(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe CapUserData
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

func TestWriteToUnmarshalPreservesDataCapUserData(t *testing.T) {
    var x, y, yUnsafe CapUserData
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

func TestSizeBytesOnTypedNilPtrCapUserData(t *testing.T) {
    var x CapUserData
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*CapUserData)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataCapUserData(t *testing.T) {
    var x, y, yUnsafe [8]CapUserData
    analysis.RandomizeValue(&x)

    size := (*CapUserData)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeCapUserDataSlice(x[:], bufUnsafe)

    UnmarshalUnsafeCapUserDataSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeCapUserDataSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroCapUserHeader(t *testing.T) {
    var x CapUserHeader
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentCapUserHeader(t *testing.T) {
    var x CapUserHeader
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataCapUserHeader(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe CapUserHeader
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

func TestWriteToUnmarshalPreservesDataCapUserHeader(t *testing.T) {
    var x, y, yUnsafe CapUserHeader
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

func TestSizeBytesOnTypedNilPtrCapUserHeader(t *testing.T) {
    var x CapUserHeader
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*CapUserHeader)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockErrCMsgIPv4(t *testing.T) {
    var x SockErrCMsgIPv4
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockErrCMsgIPv4(t *testing.T) {
    var x SockErrCMsgIPv4
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockErrCMsgIPv4(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockErrCMsgIPv4
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

func TestWriteToUnmarshalPreservesDataSockErrCMsgIPv4(t *testing.T) {
    var x, y, yUnsafe SockErrCMsgIPv4
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

func TestSizeBytesOnTypedNilPtrSockErrCMsgIPv4(t *testing.T) {
    var x SockErrCMsgIPv4
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockErrCMsgIPv4)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockErrCMsgIPv6(t *testing.T) {
    var x SockErrCMsgIPv6
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockErrCMsgIPv6(t *testing.T) {
    var x SockErrCMsgIPv6
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockErrCMsgIPv6(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockErrCMsgIPv6
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

func TestWriteToUnmarshalPreservesDataSockErrCMsgIPv6(t *testing.T) {
    var x, y, yUnsafe SockErrCMsgIPv6
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

func TestSizeBytesOnTypedNilPtrSockErrCMsgIPv6(t *testing.T) {
    var x SockErrCMsgIPv6
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockErrCMsgIPv6)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockExtendedErr(t *testing.T) {
    var x SockExtendedErr
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockExtendedErr(t *testing.T) {
    var x SockExtendedErr
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockExtendedErr(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockExtendedErr
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

func TestWriteToUnmarshalPreservesDataSockExtendedErr(t *testing.T) {
    var x, y, yUnsafe SockExtendedErr
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

func TestSizeBytesOnTypedNilPtrSockExtendedErr(t *testing.T) {
    var x SockExtendedErr
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockExtendedErr)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFOwnerEx(t *testing.T) {
    var x FOwnerEx
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFOwnerEx(t *testing.T) {
    var x FOwnerEx
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFOwnerEx(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FOwnerEx
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

func TestWriteToUnmarshalPreservesDataFOwnerEx(t *testing.T) {
    var x, y, yUnsafe FOwnerEx
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

func TestSizeBytesOnTypedNilPtrFOwnerEx(t *testing.T) {
    var x FOwnerEx
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FOwnerEx)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFlock(t *testing.T) {
    var x Flock
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFlock(t *testing.T) {
    var x Flock
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFlock(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Flock
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

func TestWriteToUnmarshalPreservesDataFlock(t *testing.T) {
    var x, y, yUnsafe Flock
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

func TestSizeBytesOnTypedNilPtrFlock(t *testing.T) {
    var x Flock
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Flock)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroStatx(t *testing.T) {
    var x Statx
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentStatx(t *testing.T) {
    var x Statx
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataStatx(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Statx
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

func TestWriteToUnmarshalPreservesDataStatx(t *testing.T) {
    var x, y, yUnsafe Statx
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

func TestSizeBytesOnTypedNilPtrStatx(t *testing.T) {
    var x Statx
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Statx)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroStatfs(t *testing.T) {
    var x Statfs
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentStatfs(t *testing.T) {
    var x Statfs
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataStatfs(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Statfs
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

func TestWriteToUnmarshalPreservesDataStatfs(t *testing.T) {
    var x, y, yUnsafe Statfs
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

func TestSizeBytesOnTypedNilPtrStatfs(t *testing.T) {
    var x Statfs
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Statfs)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEAttr(t *testing.T) {
    var x FUSEAttr
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEAttr(t *testing.T) {
    var x FUSEAttr
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEAttr(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEAttr
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

func TestWriteToUnmarshalPreservesDataFUSEAttr(t *testing.T) {
    var x, y, yUnsafe FUSEAttr
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

func TestSizeBytesOnTypedNilPtrFUSEAttr(t *testing.T) {
    var x FUSEAttr
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEAttr)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSECreateMeta(t *testing.T) {
    var x FUSECreateMeta
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSECreateMeta(t *testing.T) {
    var x FUSECreateMeta
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSECreateMeta(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSECreateMeta
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

func TestWriteToUnmarshalPreservesDataFUSECreateMeta(t *testing.T) {
    var x, y, yUnsafe FUSECreateMeta
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

func TestSizeBytesOnTypedNilPtrFUSECreateMeta(t *testing.T) {
    var x FUSECreateMeta
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSECreateMeta)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEDirentMeta(t *testing.T) {
    var x FUSEDirentMeta
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEDirentMeta(t *testing.T) {
    var x FUSEDirentMeta
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEDirentMeta(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEDirentMeta
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

func TestWriteToUnmarshalPreservesDataFUSEDirentMeta(t *testing.T) {
    var x, y, yUnsafe FUSEDirentMeta
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

func TestSizeBytesOnTypedNilPtrFUSEDirentMeta(t *testing.T) {
    var x FUSEDirentMeta
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEDirentMeta)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEEntryOut(t *testing.T) {
    var x FUSEEntryOut
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEEntryOut(t *testing.T) {
    var x FUSEEntryOut
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEEntryOut(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEEntryOut
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

func TestWriteToUnmarshalPreservesDataFUSEEntryOut(t *testing.T) {
    var x, y, yUnsafe FUSEEntryOut
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

func TestSizeBytesOnTypedNilPtrFUSEEntryOut(t *testing.T) {
    var x FUSEEntryOut
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEEntryOut)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEGetAttrIn(t *testing.T) {
    var x FUSEGetAttrIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEGetAttrIn(t *testing.T) {
    var x FUSEGetAttrIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEGetAttrIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEGetAttrIn
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

func TestWriteToUnmarshalPreservesDataFUSEGetAttrIn(t *testing.T) {
    var x, y, yUnsafe FUSEGetAttrIn
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

func TestSizeBytesOnTypedNilPtrFUSEGetAttrIn(t *testing.T) {
    var x FUSEGetAttrIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEGetAttrIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEGetAttrOut(t *testing.T) {
    var x FUSEGetAttrOut
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEGetAttrOut(t *testing.T) {
    var x FUSEGetAttrOut
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEGetAttrOut(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEGetAttrOut
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

func TestWriteToUnmarshalPreservesDataFUSEGetAttrOut(t *testing.T) {
    var x, y, yUnsafe FUSEGetAttrOut
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

func TestSizeBytesOnTypedNilPtrFUSEGetAttrOut(t *testing.T) {
    var x FUSEGetAttrOut
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEGetAttrOut)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEHeaderIn(t *testing.T) {
    var x FUSEHeaderIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEHeaderIn(t *testing.T) {
    var x FUSEHeaderIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEHeaderIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEHeaderIn
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

func TestWriteToUnmarshalPreservesDataFUSEHeaderIn(t *testing.T) {
    var x, y, yUnsafe FUSEHeaderIn
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

func TestSizeBytesOnTypedNilPtrFUSEHeaderIn(t *testing.T) {
    var x FUSEHeaderIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEHeaderIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEHeaderOut(t *testing.T) {
    var x FUSEHeaderOut
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEHeaderOut(t *testing.T) {
    var x FUSEHeaderOut
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEHeaderOut(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEHeaderOut
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

func TestWriteToUnmarshalPreservesDataFUSEHeaderOut(t *testing.T) {
    var x, y, yUnsafe FUSEHeaderOut
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

func TestSizeBytesOnTypedNilPtrFUSEHeaderOut(t *testing.T) {
    var x FUSEHeaderOut
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEHeaderOut)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEInitIn(t *testing.T) {
    var x FUSEInitIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEInitIn(t *testing.T) {
    var x FUSEInitIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEInitIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEInitIn
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

func TestWriteToUnmarshalPreservesDataFUSEInitIn(t *testing.T) {
    var x, y, yUnsafe FUSEInitIn
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

func TestSizeBytesOnTypedNilPtrFUSEInitIn(t *testing.T) {
    var x FUSEInitIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEInitIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEInitOut(t *testing.T) {
    var x FUSEInitOut
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEInitOut(t *testing.T) {
    var x FUSEInitOut
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEInitOut(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEInitOut
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

func TestWriteToUnmarshalPreservesDataFUSEInitOut(t *testing.T) {
    var x, y, yUnsafe FUSEInitOut
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

func TestSizeBytesOnTypedNilPtrFUSEInitOut(t *testing.T) {
    var x FUSEInitOut
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEInitOut)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEMkdirMeta(t *testing.T) {
    var x FUSEMkdirMeta
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEMkdirMeta(t *testing.T) {
    var x FUSEMkdirMeta
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEMkdirMeta(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEMkdirMeta
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

func TestWriteToUnmarshalPreservesDataFUSEMkdirMeta(t *testing.T) {
    var x, y, yUnsafe FUSEMkdirMeta
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

func TestSizeBytesOnTypedNilPtrFUSEMkdirMeta(t *testing.T) {
    var x FUSEMkdirMeta
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEMkdirMeta)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEMknodMeta(t *testing.T) {
    var x FUSEMknodMeta
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEMknodMeta(t *testing.T) {
    var x FUSEMknodMeta
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEMknodMeta(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEMknodMeta
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

func TestWriteToUnmarshalPreservesDataFUSEMknodMeta(t *testing.T) {
    var x, y, yUnsafe FUSEMknodMeta
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

func TestSizeBytesOnTypedNilPtrFUSEMknodMeta(t *testing.T) {
    var x FUSEMknodMeta
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEMknodMeta)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEOpID(t *testing.T) {
    var x FUSEOpID
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEOpID(t *testing.T) {
    var x FUSEOpID
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEOpID(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEOpID
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

func TestWriteToUnmarshalPreservesDataFUSEOpID(t *testing.T) {
    var x, y, yUnsafe FUSEOpID
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

func TestSizeBytesOnTypedNilPtrFUSEOpID(t *testing.T) {
    var x FUSEOpID
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEOpID)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEOpcode(t *testing.T) {
    var x FUSEOpcode
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEOpcode(t *testing.T) {
    var x FUSEOpcode
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEOpcode(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEOpcode
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

func TestWriteToUnmarshalPreservesDataFUSEOpcode(t *testing.T) {
    var x, y, yUnsafe FUSEOpcode
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

func TestSizeBytesOnTypedNilPtrFUSEOpcode(t *testing.T) {
    var x FUSEOpcode
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEOpcode)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEOpenIn(t *testing.T) {
    var x FUSEOpenIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEOpenIn(t *testing.T) {
    var x FUSEOpenIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEOpenIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEOpenIn
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

func TestWriteToUnmarshalPreservesDataFUSEOpenIn(t *testing.T) {
    var x, y, yUnsafe FUSEOpenIn
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

func TestSizeBytesOnTypedNilPtrFUSEOpenIn(t *testing.T) {
    var x FUSEOpenIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEOpenIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEOpenOut(t *testing.T) {
    var x FUSEOpenOut
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEOpenOut(t *testing.T) {
    var x FUSEOpenOut
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEOpenOut(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEOpenOut
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

func TestWriteToUnmarshalPreservesDataFUSEOpenOut(t *testing.T) {
    var x, y, yUnsafe FUSEOpenOut
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

func TestSizeBytesOnTypedNilPtrFUSEOpenOut(t *testing.T) {
    var x FUSEOpenOut
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEOpenOut)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEReadIn(t *testing.T) {
    var x FUSEReadIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEReadIn(t *testing.T) {
    var x FUSEReadIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEReadIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEReadIn
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

func TestWriteToUnmarshalPreservesDataFUSEReadIn(t *testing.T) {
    var x, y, yUnsafe FUSEReadIn
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

func TestSizeBytesOnTypedNilPtrFUSEReadIn(t *testing.T) {
    var x FUSEReadIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEReadIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEReleaseIn(t *testing.T) {
    var x FUSEReleaseIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEReleaseIn(t *testing.T) {
    var x FUSEReleaseIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEReleaseIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEReleaseIn
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

func TestWriteToUnmarshalPreservesDataFUSEReleaseIn(t *testing.T) {
    var x, y, yUnsafe FUSEReleaseIn
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

func TestSizeBytesOnTypedNilPtrFUSEReleaseIn(t *testing.T) {
    var x FUSEReleaseIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEReleaseIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSESetAttrIn(t *testing.T) {
    var x FUSESetAttrIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSESetAttrIn(t *testing.T) {
    var x FUSESetAttrIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSESetAttrIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSESetAttrIn
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

func TestWriteToUnmarshalPreservesDataFUSESetAttrIn(t *testing.T) {
    var x, y, yUnsafe FUSESetAttrIn
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

func TestSizeBytesOnTypedNilPtrFUSESetAttrIn(t *testing.T) {
    var x FUSESetAttrIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSESetAttrIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEWriteIn(t *testing.T) {
    var x FUSEWriteIn
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEWriteIn(t *testing.T) {
    var x FUSEWriteIn
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEWriteIn(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEWriteIn
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

func TestWriteToUnmarshalPreservesDataFUSEWriteIn(t *testing.T) {
    var x, y, yUnsafe FUSEWriteIn
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

func TestSizeBytesOnTypedNilPtrFUSEWriteIn(t *testing.T) {
    var x FUSEWriteIn
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEWriteIn)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroFUSEWriteOut(t *testing.T) {
    var x FUSEWriteOut
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentFUSEWriteOut(t *testing.T) {
    var x FUSEWriteOut
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataFUSEWriteOut(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe FUSEWriteOut
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

func TestWriteToUnmarshalPreservesDataFUSEWriteOut(t *testing.T) {
    var x, y, yUnsafe FUSEWriteOut
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

func TestSizeBytesOnTypedNilPtrFUSEWriteOut(t *testing.T) {
    var x FUSEWriteOut
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*FUSEWriteOut)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroRobustListHead(t *testing.T) {
    var x RobustListHead
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentRobustListHead(t *testing.T) {
    var x RobustListHead
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataRobustListHead(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe RobustListHead
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

func TestWriteToUnmarshalPreservesDataRobustListHead(t *testing.T) {
    var x, y, yUnsafe RobustListHead
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

func TestSizeBytesOnTypedNilPtrRobustListHead(t *testing.T) {
    var x RobustListHead
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*RobustListHead)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroDigestMetadata(t *testing.T) {
    var x DigestMetadata
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentDigestMetadata(t *testing.T) {
    var x DigestMetadata
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataDigestMetadata(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe DigestMetadata
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

func TestWriteToUnmarshalPreservesDataDigestMetadata(t *testing.T) {
    var x, y, yUnsafe DigestMetadata
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

func TestSizeBytesOnTypedNilPtrDigestMetadata(t *testing.T) {
    var x DigestMetadata
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*DigestMetadata)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIPCPerm(t *testing.T) {
    var x IPCPerm
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIPCPerm(t *testing.T) {
    var x IPCPerm
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIPCPerm(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IPCPerm
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

func TestWriteToUnmarshalPreservesDataIPCPerm(t *testing.T) {
    var x, y, yUnsafe IPCPerm
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

func TestSizeBytesOnTypedNilPtrIPCPerm(t *testing.T) {
    var x IPCPerm
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IPCPerm)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSysinfo(t *testing.T) {
    var x Sysinfo
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSysinfo(t *testing.T) {
    var x Sysinfo
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSysinfo(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Sysinfo
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

func TestWriteToUnmarshalPreservesDataSysinfo(t *testing.T) {
    var x, y, yUnsafe Sysinfo
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

func TestSizeBytesOnTypedNilPtrSysinfo(t *testing.T) {
    var x Sysinfo
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Sysinfo)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroNumaPolicy(t *testing.T) {
    var x NumaPolicy
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentNumaPolicy(t *testing.T) {
    var x NumaPolicy
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataNumaPolicy(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe NumaPolicy
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

func TestWriteToUnmarshalPreservesDataNumaPolicy(t *testing.T) {
    var x, y, yUnsafe NumaPolicy
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

func TestSizeBytesOnTypedNilPtrNumaPolicy(t *testing.T) {
    var x NumaPolicy
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*NumaPolicy)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIFConf(t *testing.T) {
    var x IFConf
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIFConf(t *testing.T) {
    var x IFConf
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIFConf(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IFConf
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

func TestWriteToUnmarshalPreservesDataIFConf(t *testing.T) {
    var x, y, yUnsafe IFConf
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

func TestSizeBytesOnTypedNilPtrIFConf(t *testing.T) {
    var x IFConf
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IFConf)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIFReq(t *testing.T) {
    var x IFReq
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIFReq(t *testing.T) {
    var x IFReq
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIFReq(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IFReq
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

func TestWriteToUnmarshalPreservesDataIFReq(t *testing.T) {
    var x, y, yUnsafe IFReq
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

func TestSizeBytesOnTypedNilPtrIFReq(t *testing.T) {
    var x IFReq
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IFReq)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroExtensionName(t *testing.T) {
    var x ExtensionName
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentExtensionName(t *testing.T) {
    var x ExtensionName
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataExtensionName(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ExtensionName
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

func TestWriteToUnmarshalPreservesDataExtensionName(t *testing.T) {
    var x, y, yUnsafe ExtensionName
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

func TestSizeBytesOnTypedNilPtrExtensionName(t *testing.T) {
    var x ExtensionName
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ExtensionName)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIPTEntry(t *testing.T) {
    var x IPTEntry
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIPTEntry(t *testing.T) {
    var x IPTEntry
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIPTEntry(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IPTEntry
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

func TestWriteToUnmarshalPreservesDataIPTEntry(t *testing.T) {
    var x, y, yUnsafe IPTEntry
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

func TestSizeBytesOnTypedNilPtrIPTEntry(t *testing.T) {
    var x IPTEntry
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IPTEntry)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIPTGetEntries(t *testing.T) {
    var x IPTGetEntries
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIPTGetEntries(t *testing.T) {
    var x IPTGetEntries
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIPTGetEntries(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IPTGetEntries
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

func TestWriteToUnmarshalPreservesDataIPTGetEntries(t *testing.T) {
    var x, y, yUnsafe IPTGetEntries
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

func TestSizeBytesOnTypedNilPtrIPTGetEntries(t *testing.T) {
    var x IPTGetEntries
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IPTGetEntries)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIPTGetinfo(t *testing.T) {
    var x IPTGetinfo
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIPTGetinfo(t *testing.T) {
    var x IPTGetinfo
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIPTGetinfo(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IPTGetinfo
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

func TestWriteToUnmarshalPreservesDataIPTGetinfo(t *testing.T) {
    var x, y, yUnsafe IPTGetinfo
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

func TestSizeBytesOnTypedNilPtrIPTGetinfo(t *testing.T) {
    var x IPTGetinfo
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IPTGetinfo)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIPTIP(t *testing.T) {
    var x IPTIP
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIPTIP(t *testing.T) {
    var x IPTIP
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIPTIP(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IPTIP
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

func TestWriteToUnmarshalPreservesDataIPTIP(t *testing.T) {
    var x, y, yUnsafe IPTIP
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

func TestSizeBytesOnTypedNilPtrIPTIP(t *testing.T) {
    var x IPTIP
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IPTIP)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroTableName(t *testing.T) {
    var x TableName
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTableName(t *testing.T) {
    var x TableName
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTableName(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe TableName
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

func TestWriteToUnmarshalPreservesDataTableName(t *testing.T) {
    var x, y, yUnsafe TableName
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

func TestSizeBytesOnTypedNilPtrTableName(t *testing.T) {
    var x TableName
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*TableName)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroXTCounters(t *testing.T) {
    var x XTCounters
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentXTCounters(t *testing.T) {
    var x XTCounters
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataXTCounters(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe XTCounters
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

func TestWriteToUnmarshalPreservesDataXTCounters(t *testing.T) {
    var x, y, yUnsafe XTCounters
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

func TestSizeBytesOnTypedNilPtrXTCounters(t *testing.T) {
    var x XTCounters
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*XTCounters)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroXTGetRevision(t *testing.T) {
    var x XTGetRevision
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentXTGetRevision(t *testing.T) {
    var x XTGetRevision
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataXTGetRevision(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe XTGetRevision
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

func TestWriteToUnmarshalPreservesDataXTGetRevision(t *testing.T) {
    var x, y, yUnsafe XTGetRevision
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

func TestSizeBytesOnTypedNilPtrXTGetRevision(t *testing.T) {
    var x XTGetRevision
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*XTGetRevision)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIP6TEntry(t *testing.T) {
    var x IP6TEntry
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIP6TEntry(t *testing.T) {
    var x IP6TEntry
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIP6TEntry(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IP6TEntry
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

func TestWriteToUnmarshalPreservesDataIP6TEntry(t *testing.T) {
    var x, y, yUnsafe IP6TEntry
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

func TestSizeBytesOnTypedNilPtrIP6TEntry(t *testing.T) {
    var x IP6TEntry
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IP6TEntry)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIP6TIP(t *testing.T) {
    var x IP6TIP
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIP6TIP(t *testing.T) {
    var x IP6TIP
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIP6TIP(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IP6TIP
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

func TestWriteToUnmarshalPreservesDataIP6TIP(t *testing.T) {
    var x, y, yUnsafe IP6TIP
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

func TestSizeBytesOnTypedNilPtrIP6TIP(t *testing.T) {
    var x IP6TIP
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IP6TIP)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroIP6TReplace(t *testing.T) {
    var x IP6TReplace
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentIP6TReplace(t *testing.T) {
    var x IP6TReplace
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataIP6TReplace(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe IP6TReplace
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

func TestWriteToUnmarshalPreservesDataIP6TReplace(t *testing.T) {
    var x, y, yUnsafe IP6TReplace
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

func TestSizeBytesOnTypedNilPtrIP6TReplace(t *testing.T) {
    var x IP6TReplace
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*IP6TReplace)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockAddrNetlink(t *testing.T) {
    var x SockAddrNetlink
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockAddrNetlink(t *testing.T) {
    var x SockAddrNetlink
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockAddrNetlink(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockAddrNetlink
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

func TestWriteToUnmarshalPreservesDataSockAddrNetlink(t *testing.T) {
    var x, y, yUnsafe SockAddrNetlink
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

func TestSizeBytesOnTypedNilPtrSockAddrNetlink(t *testing.T) {
    var x SockAddrNetlink
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockAddrNetlink)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroPollFD(t *testing.T) {
    var x PollFD
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentPollFD(t *testing.T) {
    var x PollFD
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataPollFD(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe PollFD
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

func TestWriteToUnmarshalPreservesDataPollFD(t *testing.T) {
    var x, y, yUnsafe PollFD
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

func TestSizeBytesOnTypedNilPtrPollFD(t *testing.T) {
    var x PollFD
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*PollFD)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataPollFD(t *testing.T) {
    var x, y, yUnsafe [8]PollFD
    analysis.RandomizeValue(&x)

    size := (*PollFD)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafePollFDSlice(x[:], bufUnsafe)

    UnmarshalUnsafePollFDSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafePollFDSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroRSeqCriticalSection(t *testing.T) {
    var x RSeqCriticalSection
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentRSeqCriticalSection(t *testing.T) {
    var x RSeqCriticalSection
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataRSeqCriticalSection(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe RSeqCriticalSection
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

func TestWriteToUnmarshalPreservesDataRSeqCriticalSection(t *testing.T) {
    var x, y, yUnsafe RSeqCriticalSection
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

func TestSizeBytesOnTypedNilPtrRSeqCriticalSection(t *testing.T) {
    var x RSeqCriticalSection
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*RSeqCriticalSection)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroRusage(t *testing.T) {
    var x Rusage
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentRusage(t *testing.T) {
    var x Rusage
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataRusage(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Rusage
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

func TestWriteToUnmarshalPreservesDataRusage(t *testing.T) {
    var x, y, yUnsafe Rusage
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

func TestSizeBytesOnTypedNilPtrRusage(t *testing.T) {
    var x Rusage
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Rusage)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSeccompData(t *testing.T) {
    var x SeccompData
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSeccompData(t *testing.T) {
    var x SeccompData
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSeccompData(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SeccompData
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

func TestWriteToUnmarshalPreservesDataSeccompData(t *testing.T) {
    var x, y, yUnsafe SeccompData
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

func TestSizeBytesOnTypedNilPtrSeccompData(t *testing.T) {
    var x SeccompData
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SeccompData)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSemInfo(t *testing.T) {
    var x SemInfo
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSemInfo(t *testing.T) {
    var x SemInfo
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSemInfo(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SemInfo
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

func TestWriteToUnmarshalPreservesDataSemInfo(t *testing.T) {
    var x, y, yUnsafe SemInfo
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

func TestSizeBytesOnTypedNilPtrSemInfo(t *testing.T) {
    var x SemInfo
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SemInfo)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSembuf(t *testing.T) {
    var x Sembuf
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSembuf(t *testing.T) {
    var x Sembuf
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSembuf(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Sembuf
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

func TestWriteToUnmarshalPreservesDataSembuf(t *testing.T) {
    var x, y, yUnsafe Sembuf
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

func TestSizeBytesOnTypedNilPtrSembuf(t *testing.T) {
    var x Sembuf
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Sembuf)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataSembuf(t *testing.T) {
    var x, y, yUnsafe [8]Sembuf
    analysis.RandomizeValue(&x)

    size := (*Sembuf)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeSembufSlice(x[:], bufUnsafe)

    UnmarshalUnsafeSembufSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeSembufSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroShmInfo(t *testing.T) {
    var x ShmInfo
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentShmInfo(t *testing.T) {
    var x ShmInfo
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataShmInfo(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ShmInfo
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

func TestWriteToUnmarshalPreservesDataShmInfo(t *testing.T) {
    var x, y, yUnsafe ShmInfo
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

func TestSizeBytesOnTypedNilPtrShmInfo(t *testing.T) {
    var x ShmInfo
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ShmInfo)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroShmParams(t *testing.T) {
    var x ShmParams
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentShmParams(t *testing.T) {
    var x ShmParams
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataShmParams(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ShmParams
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

func TestWriteToUnmarshalPreservesDataShmParams(t *testing.T) {
    var x, y, yUnsafe ShmParams
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

func TestSizeBytesOnTypedNilPtrShmParams(t *testing.T) {
    var x ShmParams
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ShmParams)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroShmidDS(t *testing.T) {
    var x ShmidDS
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentShmidDS(t *testing.T) {
    var x ShmidDS
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataShmidDS(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ShmidDS
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

func TestWriteToUnmarshalPreservesDataShmidDS(t *testing.T) {
    var x, y, yUnsafe ShmidDS
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

func TestSizeBytesOnTypedNilPtrShmidDS(t *testing.T) {
    var x ShmidDS
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ShmidDS)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSigevent(t *testing.T) {
    var x Sigevent
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSigevent(t *testing.T) {
    var x Sigevent
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSigevent(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Sigevent
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

func TestWriteToUnmarshalPreservesDataSigevent(t *testing.T) {
    var x, y, yUnsafe Sigevent
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

func TestSizeBytesOnTypedNilPtrSigevent(t *testing.T) {
    var x Sigevent
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Sigevent)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSignalSet(t *testing.T) {
    var x SignalSet
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSignalSet(t *testing.T) {
    var x SignalSet
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSignalSet(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SignalSet
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

func TestWriteToUnmarshalPreservesDataSignalSet(t *testing.T) {
    var x, y, yUnsafe SignalSet
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

func TestSizeBytesOnTypedNilPtrSignalSet(t *testing.T) {
    var x SignalSet
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SignalSet)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSignalfdSiginfo(t *testing.T) {
    var x SignalfdSiginfo
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSignalfdSiginfo(t *testing.T) {
    var x SignalfdSiginfo
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSignalfdSiginfo(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SignalfdSiginfo
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

func TestWriteToUnmarshalPreservesDataSignalfdSiginfo(t *testing.T) {
    var x, y, yUnsafe SignalfdSiginfo
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

func TestSizeBytesOnTypedNilPtrSignalfdSiginfo(t *testing.T) {
    var x SignalfdSiginfo
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SignalfdSiginfo)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroControlMessageCredentials(t *testing.T) {
    var x ControlMessageCredentials
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentControlMessageCredentials(t *testing.T) {
    var x ControlMessageCredentials
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataControlMessageCredentials(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ControlMessageCredentials
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

func TestWriteToUnmarshalPreservesDataControlMessageCredentials(t *testing.T) {
    var x, y, yUnsafe ControlMessageCredentials
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

func TestSizeBytesOnTypedNilPtrControlMessageCredentials(t *testing.T) {
    var x ControlMessageCredentials
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ControlMessageCredentials)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroInet6Addr(t *testing.T) {
    var x Inet6Addr
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentInet6Addr(t *testing.T) {
    var x Inet6Addr
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataInet6Addr(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Inet6Addr
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

func TestWriteToUnmarshalPreservesDataInet6Addr(t *testing.T) {
    var x, y, yUnsafe Inet6Addr
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

func TestSizeBytesOnTypedNilPtrInet6Addr(t *testing.T) {
    var x Inet6Addr
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Inet6Addr)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroInetAddr(t *testing.T) {
    var x InetAddr
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentInetAddr(t *testing.T) {
    var x InetAddr
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataInetAddr(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe InetAddr
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

func TestWriteToUnmarshalPreservesDataInetAddr(t *testing.T) {
    var x, y, yUnsafe InetAddr
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

func TestSizeBytesOnTypedNilPtrInetAddr(t *testing.T) {
    var x InetAddr
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*InetAddr)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroLinger(t *testing.T) {
    var x Linger
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentLinger(t *testing.T) {
    var x Linger
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataLinger(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Linger
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

func TestWriteToUnmarshalPreservesDataLinger(t *testing.T) {
    var x, y, yUnsafe Linger
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

func TestSizeBytesOnTypedNilPtrLinger(t *testing.T) {
    var x Linger
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Linger)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockAddrInet(t *testing.T) {
    var x SockAddrInet
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockAddrInet(t *testing.T) {
    var x SockAddrInet
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockAddrInet(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockAddrInet
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

func TestWriteToUnmarshalPreservesDataSockAddrInet(t *testing.T) {
    var x, y, yUnsafe SockAddrInet
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

func TestSizeBytesOnTypedNilPtrSockAddrInet(t *testing.T) {
    var x SockAddrInet
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockAddrInet)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockAddrInet6(t *testing.T) {
    var x SockAddrInet6
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockAddrInet6(t *testing.T) {
    var x SockAddrInet6
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockAddrInet6(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockAddrInet6
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

func TestWriteToUnmarshalPreservesDataSockAddrInet6(t *testing.T) {
    var x, y, yUnsafe SockAddrInet6
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

func TestSizeBytesOnTypedNilPtrSockAddrInet6(t *testing.T) {
    var x SockAddrInet6
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockAddrInet6)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockAddrLink(t *testing.T) {
    var x SockAddrLink
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockAddrLink(t *testing.T) {
    var x SockAddrLink
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockAddrLink(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockAddrLink
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

func TestWriteToUnmarshalPreservesDataSockAddrLink(t *testing.T) {
    var x, y, yUnsafe SockAddrLink
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

func TestSizeBytesOnTypedNilPtrSockAddrLink(t *testing.T) {
    var x SockAddrLink
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockAddrLink)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroSockAddrUnix(t *testing.T) {
    var x SockAddrUnix
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentSockAddrUnix(t *testing.T) {
    var x SockAddrUnix
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataSockAddrUnix(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe SockAddrUnix
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

func TestWriteToUnmarshalPreservesDataSockAddrUnix(t *testing.T) {
    var x, y, yUnsafe SockAddrUnix
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

func TestSizeBytesOnTypedNilPtrSockAddrUnix(t *testing.T) {
    var x SockAddrUnix
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*SockAddrUnix)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroTCPInfo(t *testing.T) {
    var x TCPInfo
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTCPInfo(t *testing.T) {
    var x TCPInfo
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTCPInfo(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe TCPInfo
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

func TestWriteToUnmarshalPreservesDataTCPInfo(t *testing.T) {
    var x, y, yUnsafe TCPInfo
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

func TestSizeBytesOnTypedNilPtrTCPInfo(t *testing.T) {
    var x TCPInfo
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*TCPInfo)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroClockT(t *testing.T) {
    var x ClockT
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentClockT(t *testing.T) {
    var x ClockT
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataClockT(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ClockT
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

func TestWriteToUnmarshalPreservesDataClockT(t *testing.T) {
    var x, y, yUnsafe ClockT
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

func TestSizeBytesOnTypedNilPtrClockT(t *testing.T) {
    var x ClockT
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ClockT)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroItimerVal(t *testing.T) {
    var x ItimerVal
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentItimerVal(t *testing.T) {
    var x ItimerVal
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataItimerVal(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe ItimerVal
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

func TestWriteToUnmarshalPreservesDataItimerVal(t *testing.T) {
    var x, y, yUnsafe ItimerVal
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

func TestSizeBytesOnTypedNilPtrItimerVal(t *testing.T) {
    var x ItimerVal
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*ItimerVal)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroItimerspec(t *testing.T) {
    var x Itimerspec
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentItimerspec(t *testing.T) {
    var x Itimerspec
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataItimerspec(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Itimerspec
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

func TestWriteToUnmarshalPreservesDataItimerspec(t *testing.T) {
    var x, y, yUnsafe Itimerspec
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

func TestSizeBytesOnTypedNilPtrItimerspec(t *testing.T) {
    var x Itimerspec
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Itimerspec)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroStatxTimestamp(t *testing.T) {
    var x StatxTimestamp
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentStatxTimestamp(t *testing.T) {
    var x StatxTimestamp
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataStatxTimestamp(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe StatxTimestamp
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

func TestWriteToUnmarshalPreservesDataStatxTimestamp(t *testing.T) {
    var x, y, yUnsafe StatxTimestamp
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

func TestSizeBytesOnTypedNilPtrStatxTimestamp(t *testing.T) {
    var x StatxTimestamp
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*StatxTimestamp)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroTimeT(t *testing.T) {
    var x TimeT
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTimeT(t *testing.T) {
    var x TimeT
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTimeT(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe TimeT
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

func TestWriteToUnmarshalPreservesDataTimeT(t *testing.T) {
    var x, y, yUnsafe TimeT
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

func TestSizeBytesOnTypedNilPtrTimeT(t *testing.T) {
    var x TimeT
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*TimeT)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroTimerID(t *testing.T) {
    var x TimerID
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTimerID(t *testing.T) {
    var x TimerID
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTimerID(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe TimerID
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

func TestWriteToUnmarshalPreservesDataTimerID(t *testing.T) {
    var x, y, yUnsafe TimerID
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

func TestSizeBytesOnTypedNilPtrTimerID(t *testing.T) {
    var x TimerID
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*TimerID)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroTimespec(t *testing.T) {
    var x Timespec
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTimespec(t *testing.T) {
    var x Timespec
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTimespec(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Timespec
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

func TestWriteToUnmarshalPreservesDataTimespec(t *testing.T) {
    var x, y, yUnsafe Timespec
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

func TestSizeBytesOnTypedNilPtrTimespec(t *testing.T) {
    var x Timespec
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Timespec)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataTimespec(t *testing.T) {
    var x, y, yUnsafe [8]Timespec
    analysis.RandomizeValue(&x)

    size := (*Timespec)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeTimespecSlice(x[:], bufUnsafe)

    UnmarshalUnsafeTimespecSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeTimespecSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroTimeval(t *testing.T) {
    var x Timeval
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTimeval(t *testing.T) {
    var x Timeval
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTimeval(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Timeval
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

func TestWriteToUnmarshalPreservesDataTimeval(t *testing.T) {
    var x, y, yUnsafe Timeval
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

func TestSizeBytesOnTypedNilPtrTimeval(t *testing.T) {
    var x Timeval
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Timeval)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSafeMarshalUnmarshalSlicePreservesDataTimeval(t *testing.T) {
    var x, y, yUnsafe [8]Timeval
    analysis.RandomizeValue(&x)

    size := (*Timeval)(nil).SizeBytes() * len(x)
    buf := bytes.NewBuffer(make([]byte, size))
    buf.Reset()
    if err := binary.Write(buf, usermem.ByteOrder, x[:]); err != nil {
        t.Fatal(fmt.Sprintf("binary.Write failed: %v", err))
    }
    bufUnsafe := make([]byte, size)
    MarshalUnsafeTimevalSlice(x[:], bufUnsafe)

    UnmarshalUnsafeTimevalSlice(y[:], buf.Bytes())
    if !reflect.DeepEqual(x, y) {
        t.Fatal(fmt.Sprintf("Data corrupted across binary.Write/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, y))
    }
    UnmarshalUnsafeTimevalSlice(yUnsafe[:], bufUnsafe)
    if !reflect.DeepEqual(x, yUnsafe) {
        t.Fatal(fmt.Sprintf("Data corrupted across MarshalUnsafeSlice/UnmarshalUnsafeSlice cycle:\nBefore: %+v\nAfter: %+v\n", x, yUnsafe))
    }

}

func TestSizeNonZeroTms(t *testing.T) {
    var x Tms
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTms(t *testing.T) {
    var x Tms
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTms(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Tms
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

func TestWriteToUnmarshalPreservesDataTms(t *testing.T) {
    var x, y, yUnsafe Tms
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

func TestSizeBytesOnTypedNilPtrTms(t *testing.T) {
    var x Tms
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Tms)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroUtime(t *testing.T) {
    var x Utime
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUtime(t *testing.T) {
    var x Utime
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUtime(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Utime
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

func TestWriteToUnmarshalPreservesDataUtime(t *testing.T) {
    var x, y, yUnsafe Utime
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

func TestSizeBytesOnTypedNilPtrUtime(t *testing.T) {
    var x Utime
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Utime)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroTermios(t *testing.T) {
    var x Termios
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentTermios(t *testing.T) {
    var x Termios
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataTermios(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Termios
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

func TestWriteToUnmarshalPreservesDataTermios(t *testing.T) {
    var x, y, yUnsafe Termios
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

func TestSizeBytesOnTypedNilPtrTermios(t *testing.T) {
    var x Termios
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Termios)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroWindowSize(t *testing.T) {
    var x WindowSize
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentWindowSize(t *testing.T) {
    var x WindowSize
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataWindowSize(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe WindowSize
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

func TestWriteToUnmarshalPreservesDataWindowSize(t *testing.T) {
    var x, y, yUnsafe WindowSize
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

func TestSizeBytesOnTypedNilPtrWindowSize(t *testing.T) {
    var x WindowSize
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*WindowSize)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroWinsize(t *testing.T) {
    var x Winsize
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentWinsize(t *testing.T) {
    var x Winsize
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataWinsize(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe Winsize
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

func TestWriteToUnmarshalPreservesDataWinsize(t *testing.T) {
    var x, y, yUnsafe Winsize
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

func TestSizeBytesOnTypedNilPtrWinsize(t *testing.T) {
    var x Winsize
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*Winsize)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

func TestSizeNonZeroUtsName(t *testing.T) {
    var x UtsName
    if x.SizeBytes() == 0 {
        t.Fatal("Marshallable.SizeBytes() should not return zero")
    }
}

func TestSuspectAlignmentUtsName(t *testing.T) {
    var x UtsName
    analysis.AlignmentCheck(t, reflect.TypeOf(x))
}

func TestSafeMarshalUnmarshalPreservesDataUtsName(t *testing.T) {
    var x, y, z, yUnsafe, zUnsafe UtsName
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

func TestWriteToUnmarshalPreservesDataUtsName(t *testing.T) {
    var x, y, yUnsafe UtsName
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

func TestSizeBytesOnTypedNilPtrUtsName(t *testing.T) {
    var x UtsName
    sizeFromConcrete := x.SizeBytes()
    sizeFromTypedNilPtr := (*UtsName)(nil).SizeBytes()

    if sizeFromTypedNilPtr != sizeFromConcrete {
        t.Fatalf("SizeBytes() on typed nil pointer (%v) doesn't match size returned by a concrete object (%v).\n", sizeFromTypedNilPtr, sizeFromConcrete)
    }
}

