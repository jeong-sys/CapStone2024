use std::fs::File;
use std::io::{self, Write};
use std::path::Path;

pub fn configure_cgroups() -> io::Result<()> {
    let cgroup_path = "/sys/fs/cgroup/memory/my_container_runtime";
    let memory_limit_path = Path::new(cgroup_path).join("memory.limit_in_bytes");

    // cgroup 디렉토리 생성
    std::fs::create_dir_all(cgroup_path)?;

    // 메모리 제한 설정 파일 열기
    let mut file = File::create(memory_limit_path)?;

    // 메모리 제한 값 쓰기 (예: 512MB)
    writeln!(file, "{}", 512 * 1024 * 1024)?;

    Ok(())
}
