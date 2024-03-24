mod namespace

fn main() {
    namespace::setup_namespace();
    cgroup::configure_cgroups();
    process::start_process();

}