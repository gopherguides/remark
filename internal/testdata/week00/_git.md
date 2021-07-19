# Installing Git

Git is a popular SCM (source code management) tool that is used heavily by Go projects.

## Installing Git on Windows

1. Navigate to [http://git-scm.com/downloads](http://git-scm.com/download/win)
1. Download the 64-bit git for windows setup installer.
1. Once downloaded Windows should prompt you to run the installer. Click `Run`.
1. Click through all the "next" prompts.

## Installing Git on MacOS - TODO

## Installing Git on Unix - TODO

## Verifying The Git Installation - TODO

1. Click the "Start" menu
1. Type in `cmd` and press "enter".
1. In the terminal type `git --version` and press "enter".
1. You should receive the following message: `git version 2.x.x.windows`

If you do not get the above message your installation is *NOT* complete. Recheck the steps and try again.

The `x`'s in the version number is a placeholder for the most recent version number.

## Configuring Git - TODO

1. Click the "Start" menu
1. Type in `cmd` and press "enter".
1. In the terminal type `git config --global user.email "<your email address>"` and press "enter".
1. In the terminal type `git config --global user.name "<your name>"` and press "enter".
1. Type `git config -l` to verify your information.

git config --global init.defaultBranch main

## Configure Git Keys/Certificates (Optional)

Git can be configured to use SSH keys. This makes it so you don't have to constantly provide your password at the command line to get or send commits to the repositories.  You can find detailed instructions on how to do this here:

[Generating a new SSH key and adding it to the ssh-agent](https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/)

For additional security, you can also sign your commits and they will get the `verified` icon on github.  Follows these instructions to set up signed commits using GPG:

[Signing commits using GPG](https://help.github.com/articles/signing-commits-using-gpg/)

