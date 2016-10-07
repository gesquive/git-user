class GitUser < Formula
  version '2.0.3'
  desc "Git plugin that allows you to save multiple user profiles and set them as project defaults"
  homepage "https://github.com/gesquive/git-user"
  url "https://github.com/gesquive/git-user/releases/download/v#{version}/git-user-v#{version}-osx-x64.tar.gz"
  sha256 "5610ec3d3f5de257227007e0422695bfa94efb301dff3faa28371d9e2cc5ce06"

  conflicts_with "git-user"

  def install
    bin.install "git-user"
    man.mkpath
    man1.install "man/git-user.1", "man/git-user_add.1", "man/git-user_del.1", "man/git-user_edit.1", "man/git-user_list.1", "man/git-user_rm.1", "man/git-user_set.1"
  end

  test do
    output = shell_output(bin/"git-user --version")
    assert_match "git-user v#{version}", output
  end
end
