use nix::sched::{unshare, CloneFlags};
use std::process::Command;

pub fn setup_namespace() {
    // UTS 네임스페이스 생성
    unshare(CloneFlags::CLONE_NEWUTS).expect("Failed to unshare UTS namespace");
    // IPC 네임스페이스 생성
    unshare(CloneFlags::CLONE_NEWIPC).expect("Failed to unshare IPC namespace");
    // PID 네임스페이스 생성
    unshare(CloneFlags::CLONE_NEWPID).expect("Failed to unshare PID namespace");
    // 네트워크 네임스페이스 생성
    unshare(CloneFlags::CLONE_NEWNET).expect("Failed to unshare NET namespace");
    // 마운트 네임스페이스 생성
    unshare(CloneFlags::CLONE_NEWS).expect("Failed to unshare MNT namespace");
}
