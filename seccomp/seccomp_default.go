package seccomp

import (
	"fmt"

	"github.com/docker/docker/api/types"
)

func arches(arch string) []types.Architecture {
	fmt.Println(arch)
	return []types.Architecture{
		{
			Arch:      types.ArchX86_64,
			SubArches: []types.Arch{types.ArchX86, types.ArchX32},
		},
		{
			Arch:      types.ArchAARCH64,
			SubArches: []types.Arch{types.ArchARM},
		},
		{
			Arch:      types.ArchMIPS64,
			SubArches: []types.Arch{types.ArchMIPS, types.ArchMIPS64N32},
		},
		{
			Arch:      types.ArchMIPS64N32,
			SubArches: []types.Arch{types.ArchMIPS, types.ArchMIPS64},
		},
		{
			Arch:      types.ArchMIPSEL64,
			SubArches: []types.Arch{types.ArchMIPSEL, types.ArchMIPSEL64N32},
		},
		{
			Arch:      types.ArchMIPSEL64N32,
			SubArches: []types.Arch{types.ArchMIPSEL, types.ArchMIPSEL64},
		},
		{
			Arch:      types.ArchS390X,
			SubArches: []types.Arch{types.ArchS390},
		},
	}
}

// DefaultProfile defines the whitelist for the default seccomp profile.
func DefaultProfile(Syscalls []string, arch string) *types.Seccomp {
	syscalls := []*types.Syscall{
		{
			Names:  Syscalls,
			Action: types.ActAllow,
			Args:   []*types.Arg{},
		},
		{
			Names:  []string{"personality"},
			Action: types.ActAllow,
			Args: []*types.Arg{
				{
					Index: 0,
					Value: 0x0,
					Op:    types.OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: types.ActAllow,
			Args: []*types.Arg{
				{
					Index: 0,
					Value: 0x0008,
					Op:    types.OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: types.ActAllow,
			Args: []*types.Arg{
				{
					Index: 0,
					Value: 0x20000,
					Op:    types.OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: types.ActAllow,
			Args: []*types.Arg{
				{
					Index: 0,
					Value: 0x20008,
					Op:    types.OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: types.ActAllow,
			Args: []*types.Arg{
				{
					Index: 0,
					Value: 0xffffffff,
					Op:    types.OpEqualTo,
				},
			},
		},
		{
			Names: []string{
				"sync_file_range2",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Arches: []string{"ppc64le"},
			},
		},
		{
			Names: []string{
				"arm_fadvise64_64",
				"arm_sync_file_range",
				"sync_file_range2",
				"breakpoint",
				"cacheflush",
				"set_tls",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Arches: []string{"arm", "arm64"},
			},
		},
		{
			Names: []string{
				"arch_prctl",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Arches: []string{"amd64", "x32"},
			},
		},
		{
			Names: []string{
				"modify_ldt",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Arches: []string{"amd64", "x32", "x86"},
			},
		},
		{
			Names: []string{
				"s390_pci_mmio_read",
				"s390_pci_mmio_write",
				"s390_runtime_instr",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Arches: []string{"s390", "s390x"},
			},
		},
		{
			Names: []string{
				"open_by_handle_at",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_DAC_READ_SEARCH"},
			},
		},
		{
			Names: []string{
				"bpf",
				"clone",
				"fanotify_init",
				"lookup_dcookie",
				"mount",
				"name_to_handle_at",
				"perf_event_open",
				"quotactl",
				"setdomainname",
				"sethostname",
				"setns",
				"syslog",
				"umount",
				"umount2",
				"unshare",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_ADMIN"},
			},
		},
		{
			Names: []string{
				"reboot",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_BOOT"},
			},
		},
		{
			Names: []string{
				"chroot",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_CHROOT"},
			},
		},
		{
			Names: []string{
				"delete_module",
				"init_module",
				"finit_module",
				"query_module",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_MODULE"},
			},
		},
		{
			Names: []string{
				"acct",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_PACCT"},
			},
		},
		{
			Names: []string{
				"kcmp",
				"process_vm_readv",
				"process_vm_writev",
				"ptrace",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_PTRACE"},
			},
		},
		{
			Names: []string{
				"iopl",
				"ioperm",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_RAWIO"},
			},
		},
		{
			Names: []string{
				"settimeofday",
				"stime",
				"clock_settime",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_TIME"},
			},
		},
		{
			Names: []string{
				"vhangup",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_TTY_CONFIG"},
			},
		},
		{
			Names: []string{
				"get_mempolicy",
				"mbind",
				"set_mempolicy",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYS_NICE"},
			},
		},
		{
			Names: []string{
				"syslog",
			},
			Action: types.ActAllow,
			Args:   []*types.Arg{},
			Includes: types.Filter{
				Caps: []string{"CAP_SYSLOG"},
			},
		},
	}

	return &types.Seccomp{
		DefaultAction: types.ActErrno,
		ArchMap:       arches(arch),
		Syscalls:      syscalls,
	}
}
