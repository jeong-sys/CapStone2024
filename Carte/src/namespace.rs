use nix::sched::{unshare, CloneFlags};
use std::process::Command;

pub fn setup_namespace() {
    // UTS 네임스페이스 생성
    unshare(CloneFlags::CLONE_NEWUTS).expect("Failed to unshare UTS namespace");

    // `hostname` 명령어를 실행하여 네임스페이스 내의 호스트 이름을 확인
    let output = Command::new("hostname")
        .output()
        .expect("Failed to execute hostname command");

    println!("UTS Namespace Hostname: {}", String::from_utf8_lossy(&output.stdout));
}