#!/bin/bash

set -e

function copy_dir() {
    DIR=${1}
    mkdir -p ${DIR}/
    cp -r bazel-gvisor/bazel-out/k8-fastbuild/bin/${DIR}_/gvisor.dev/gvisor/${DIR}/* ${DIR}/
    chmod 644 ${DIR}/*
}

bazel build //runsc //pkg/sentry/kernel/memevent //pkg/sentry/strace //test/packetimpact/proto:*

copy_dir "pkg/metric/metric_go_proto"
copy_dir "pkg/sentry/kernel/uncaught_signal_go_proto"
copy_dir "pkg/sentry/strace/strace_go_proto"
copy_dir "pkg/eventchannel/eventchannel_go_proto"
copy_dir "pkg/sentry/arch/registers_go_proto"
copy_dir "pkg/sentry/unimpl/unimplemented_syscall_go_proto"
copy_dir "pkg/sentry/kernel/memevent/memory_events_go_proto"
copy_dir "test/packetimpact/proto/posix_server_go_proto"

# mkdir -p pkg/sentry/kernel/uncaught_signal_go_proto/
# cp -r bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/kernel/uncaught_signal_go_proto_/gvisor.dev/gvisor/pkg/sentry/kernel/uncaught_signal_go_proto/* \
#     pkg/sentry/kernel/uncaught_signal_go_proto/
# chmod 644 pkg/sentry/kernel/uncaught_signal_go_proto/*

# mkdir -p pkg/sentry/strace/strace_go_proto/
# cp -r bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/strace/strace_go_proto_/gvisor.dev/gvisor/pkg/sentry/strace/strace_go_proto/* \
#     pkg/sentry/strace/strace_go_proto/
# chmod 644 pkg/sentry/strace/strace_go_proto/*

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/bits/bits64.go pkg/bits/
chmod 644 pkg/bits/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/refs/weak_ref_list.go pkg/refs/
chmod 644 pkg/refs/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/waiter/waiter_list.go pkg/waiter/
chmod 644 pkg/waiter/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/fs/lock/lock_range.go pkg/sentry/fs/lock/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/fs/lock/lock_set.go pkg/sentry/fs/lock/
chmod 644 pkg/sentry/fs/lock/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/memmap/mappable_range.go pkg/sentry/memmap/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/memmap/mapping_set_impl.go pkg/sentry/memmap/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/memmap/file_range.go pkg/sentry/memmap/
chmod 644 pkg/sentry/memmap/*.go

cp bazel-gvisor/bazel-out/host/bin/pkg/sentry/platform/ring0/pagetables/walker_*.go pkg/sentry/platform/ring0/pagetables/
chmod 644 pkg/sentry/platform/ring0/pagetables/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/state/deferred_list.go pkg/state/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/state/addr_range.go pkg/state/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/state/addr_set.go pkg/state/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/state/complete_list.go pkg/state/
chmod 644 pkg/state/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/usermem/addr_range.go pkg/usermem/
chmod 644 pkg/usermem/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/sock_err_list.go pkg/tcpip/
chmod 644 pkg/tcpip/*.go


cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/stack/linkaddrentry_list.go pkg/tcpip/stack/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/stack/neighbor_entry_list.go pkg/tcpip/stack/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/stack/packet_buffer_list.go pkg/tcpip/stack/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/stack/tuple_list.go pkg/tcpip/stack/
chmod 644 pkg/tcpip/stack/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/network/fragmentation/reassembler_list.go pkg/tcpip/network/fragmentation/
chmod 644 pkg/tcpip/network/fragmentation/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/packet/packet_list.go pkg/tcpip/transport/packet/
chmod 644 pkg/tcpip/transport/packet/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/raw/raw_packet_list.go pkg/tcpip/transport/raw/
chmod 644 pkg/tcpip/transport/raw/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/tcp/tcp_endpoint_list.go pkg/tcpip/transport/tcp/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/tcp/tcp_segment_list.go pkg/tcpip/transport/tcp/
chmod 644 pkg/tcpip/transport/tcp/*.go


cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/udp/udp_packet_list.go pkg/tcpip/transport/udp/
chmod 644 pkg/tcpip/transport/udp/*.go

cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/icmp/icmp_packet_list.go pkg/tcpip/transport/icmp/
chmod 644 pkg/tcpip/transport/icmp/*.go


cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/tcpip/transport/icmp/icmp_packet_list.go pkg/tcpip/transport/icmp/
chmod 644 pkg/tcpip/transport/icmp/*.go



for d in pkg/marshal/primitive pkg/abi/linux pkg/sentry/kernel/auth; do
  cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/${d}/*_abi_autogen_*.go ${d}
  chmod 644 ${d}/*.go
done

rsync -avzh --include='*_set.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_set_impl.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_autogen.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_range.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_list.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_set.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_abi_autogen_*.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_refs.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='fstree.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_unsafe.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_impl_amd64.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_impl_arm64.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='bits32.go' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/
rsync -avzh --include='*_impl_amd64.s' --include='*/' --exclude='*' bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/ pkg/

mkdir -p pkg/sentry/loader/vdsodata/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/loader/vdsodata/vdso_arm64.go pkg/sentry/loader/vdsodata/
cp bazel-gvisor/bazel-out/k8-fastbuild-ST-4c64f0b3d5c7/bin/pkg/sentry/loader/vdsodata/vdso_amd64.go pkg/sentry/loader/vdsodata/

find pkg/ -name "*.go" -exec chmod 644 {} +
