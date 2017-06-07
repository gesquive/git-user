class GitUser < Formula
  version '2.0.5'
  desc "Git plugin that allows you to save multiple user profiles and set them as project defaults"
  homepage "https://github.com/gesquive/git-user"
  url "https://github.com/gesquive/git-user/releases/download/v#{version}/git-user-v#{version}-osx-x64.tar.gz"
  sha256 "afe38cd5f6e7d93f54f7a85b8ba0e4fec4d68d9a3d1b74c53cfda3b83de7bfd5"

  conflicts_with "git-user"

  def install
    bin.install "git-user"
    man.mkpath
    man1.install "man/git-user.1", "man/git-user-add.1", "man/git-user-del.1", "man/git-user-edit.1", "man/git-user-list.1", "man/git-user-rm.1", "man/git-user-set.1"
  end

  test do
    output = shell_output(bin/"git-user --version")
    assert_match "git-user v#{version}", output
  end
end
