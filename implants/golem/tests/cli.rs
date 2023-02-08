use assert_cmd::prelude::*; // Add methods on commands
use predicates::prelude::*; // Used for writing assertions
use std::process::{Command,Stdio}; // Run programs
use std::io::prelude::*;
use std::str;

// Test running `./golem ./nonexistentdir/run.tome`
#[test]
fn test_golem_main_file_not_found() -> anyhow::Result<()> {
    let mut cmd = Command::cargo_bin("golem")?;

    cmd.arg("nonexistentdir/run.tome");
    cmd.assert()
        .failure()
        .stderr(predicate::str::contains("[TASK ERROR] nonexistentdir/run.tome: No such file or directory (os error 2)"));

    Ok(())
}

// Test running `./golem ./working_dir/tomes/syntax_fail.tome`
#[test]
fn test_golem_main_syntax_fail() -> anyhow::Result<()> {
    let mut cmd = Command::cargo_bin("golem")?;

    cmd.arg("working_dir/tomes/syntax_fail.tome");
    cmd.assert()
        .failure()
        .stderr(predicate::str::contains("[TASK ERROR] working_dir/tomes/syntax_fail.tome: error: Parse error: unexpected string literal 'win' here"));

    Ok(())
}


// Test running `./golem ./working_dir/tomes/hello_world.tome`
#[test]
fn test_golem_main_basic_non_interactive() -> anyhow::Result<()> {
    let mut cmd = Command::cargo_bin("golem")?;

    cmd.arg("working_dir/tomes/hello_world.tome");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains("[\"HELLO\"]"));

    Ok(())
}


// Test running `./golem ./working_dir/tomes/eldritch_test.tome`
#[test]
fn test_golem_main_basic_eldritch_non_interactive() -> anyhow::Result<()> {
    let mut cmd = Command::cargo_bin("golem")?;

    cmd.arg("working_dir/tomes/eldritch_test.tome");
    cmd.assert()
        .success()
        .stdout(predicate::str::contains("[\"[\\\"append\\\", \\\"copy\\\", \\\"download\\\", \\\"exists\\\", \\\"hash\\\", \\\"is_dir\\\", \\\"is_file\\\", \\\"mkdir\\\", \\\"read\\\", \\\"remove\\\", \\\"rename\\\", \\\"replace\\\", \\\"replace_all\\\", \\\"timestomp\\\", \\\"write\\\"]\"]"));

    Ok(())
}

// Test running `echo -e "test_var = 'hello'\nprint(test_var)" | ./golem` for interactive mode.
// verifies that the process exits successfully. Not the output of the command.
// The way the interactive context returns data doesn't seem to work with how Command::stdout() works.
#[test]
fn test_golem_main_basic_interactive() -> anyhow::Result<()> {
    let golem_exec = Command::cargo_bin("golem")?;
    let golem_exec_path = golem_exec.get_program();

    let mut child = Command::new(golem_exec_path)
        .stdin(Stdio::piped())
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .spawn()
        .expect("Failed to spawn child process");

    let mut stdin = child.stdin.take().expect("Failed to open stdin");
    std::thread::spawn(move || {
        let _ = stdin.write_all("test_var = 'hello'\nprint(test_var)".as_bytes());
    });

    let output = child.wait_with_output().expect("Failed to read stdout");
    assert_eq!(str::from_utf8(&output.stderr)?, "hello\n");

    Ok(())
}