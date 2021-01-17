load("//tools:defs.bzl", "arch_genrule", "go_library")
load("//tools/go_generics:defs.bzl", "go_template", "go_template_instance")

package(licenses = ["notice"])

go_template(
    name = "defs_amd64",
    srcs = [
        "defs.gotemplate",
        "defs_amd64.gotemplate",
        "offsets_amd64.gotemplate",
        "x86.gotemplate",
    ],
    visibility = [":__subpackages__"],
)

go_template(
    name = "defs_arm64",
    srcs = [
        "aarch64.gotemplate",
        "defs.gotemplate",
        "defs_arm64.gotemplate",
        "offsets_arm64.gotemplate",
    ],
    visibility = [":__subpackages__"],
)

go_template_instance(
    name = "defs_impl_amd64",
    out = "defs_impl_amd64.go",
    package = "ring0",
    template = ":defs_amd64",
)

go_template_instance(
    name = "defs_impl_arm64",
    out = "defs_impl_arm64.go",
    package = "ring0",
    template = ":defs_arm64",
)

arch_genrule(
    name = "entry_impl_amd64",
    srcs = ["entry_amd64.stemplate"],
    outs = ["entry_impl_amd64.s"],
    cmd = "(echo -e '// build +amd64\\n' && QEMU $(location //pkg/sentry/platform/ring0/gen_offsets) && cat $(location entry_amd64.stemplate)) > $@",
    tools = ["//pkg/sentry/platform/ring0/gen_offsets"],
)

arch_genrule(
    name = "entry_impl_arm64",
    srcs = ["entry_arm64.s"],
    outs = ["entry_impl_arm64.s"],
    cmd = "(echo -e '// build +arm64\\n' && QEMU $(location //pkg/sentry/platform/ring0/gen_offsets) && cat $(location entry_arm64.s)) > $@",
    tools = ["//pkg/sentry/platform/ring0/gen_offsets"],
)

go_library(
    name = "ring0",
    srcs = [
        "defs_impl_amd64.go",
        "defs_impl_arm64.go",
        "entry_amd64.go",
        "entry_arm64.go",
        "entry_impl_amd64.s",
        "entry_impl_arm64.s",
        "kernel.go",
        "kernel_amd64.go",
        "kernel_arm64.go",
        "kernel_unsafe.go",
        "lib_amd64.go",
        "lib_amd64.s",
        "lib_arm64.go",
        "lib_arm64.s",
        "ring0.go",
    ],
    visibility = ["//pkg/sentry:internal"],
    deps = [
        "//pkg/cpuid",
        "//pkg/safecopy",
        "//pkg/sentry/arch",
        "//pkg/sentry/platform/ring0/pagetables",
        "//pkg/usermem",
    ],
)
