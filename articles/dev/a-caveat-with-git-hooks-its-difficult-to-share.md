---- Jul 13 13:50:39 IST 2023
## A caveat with git hooks: it's difficult to share them via version control

### Sharing git hooks between team members is slightly cumbersome. You can't just commit the contents of your ".git/hooks" directory, and walkarounds are flaky.

#### Git ignores the ".git" directory (and therefore everything inside ".git/hooks" as well) and there's not much that you can do about it.

Some solutions to this are provided by commit managers like Husky. But it's a long story and another kind of complexity. Instead this article discusses more native solutions which imply as low maintenance cost as possible.

Therefore if you want to share your hooks with other team members you will need to keep them outside .git directory. (must be executable, so make sure you chmod +x) This will allow you to commit your hooks... But git will stop running them on commit. There are at least two solutions for that:

1) you can instruct your teammates to manually copy the hooks to .git folder when cloning the repo. (And fetching too! Because hooks will change). And one day someone will forget to do that.

2) A more clever solution: you can specify git config hooksPath, pointing to your hooks that live outside .git directory. (Show how) Now hooks will be run... But your git config is local to your machine and doesn't travel with your repo. Therefore your teammates will have to manually specify the hooksPath when cloning your repo.
But at least this time your teammates will have to do manual stuff only once when cloning. Later changes to hooks won't require any action and that's a little better.
So this solution is preferable though anyway some of your colleagues will forget this too.
In case of a Node project this can be automated:
    https://stackoverflow.com/a/55958779
Config can be set automatically with makefile:
    https://www.viget.com/articles/two-ways-to-share-git-hooks-with-your-team/


3) Quote from stackOv: "Theoretically, you could create a hooks directory (or whatever name you prefer) in your project directory with all the scripts, and then symlink them in .git/hooks. Of course, each person who cloned the repo would have to set up these symlinks (although you could get really fancy and have a deploy script that the cloner could run to set them up semi-automatically)."

4) Another kind of solution could be to keep your hooks in a separate repo. This approach is encouraged by a library pre-commit for example. This again makes your hooks easy to share but difficult to make git actually run them  when necessary.




https://stackoverflow.com/questions/427207/can-git-hook-scripts-be-managed-along-with-the-repository

