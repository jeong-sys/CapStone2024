use std::process::{self, Command};

pub fn start_process() -> process::ExitStatus {
    // 새 프로세스(예: `/bin/sh`) 실행
    let child = Command::new("/bin/sh")
        .arg("-c")
        .arg("echo Hello from the container!")
        .spawn()
        .expect("Failed to start process");

    // 자식 프로세스의 종료를 기다림
    let output = child
        .wait_with_output()
        .expect("Failed to wait on child process");

    output.status
}
