---- Jul 13 13:50:39 IST 2023
# A caveat with git hooks: it's difficult to share them via version control

## You can't just commit the contents of your ".git/hooks" directory. To ensure that your colleagues use the same hooks you need some walkarounds, and all of them are a bit flaky.

Git ignores the ".git" directory and therefore everything inside ".git/hooks" as well. Some solutions to this are provided by commit managers like Husky. But it's a long story and another kind of complexity. Instead this article explores more native solutions which imply as low maintenance cost as possible.

__Solution 1.__ You can instruct your teammates to manually copy all hooks to .git folder when cloning the repo. (And fetching too! Because hooks will change). And one day someone will forget to do that. So this solution is not suitable for real life. 

You have no choice but to keep your hooks outside .git directory if you want to share them with teammates. This will allow you to commit your hooks. But git will refuse to run them unless you implement one of the following solutions.

__Solution 2.__ You can use a hooksPath config variable. Set it to a directory where you keep your hooks and git will start running them on commit:

```
git config core.hooksPath hooks_dir
```

That's easy. But your git config is __local__ to your machine and doesn't travel with your repo. Your teammates need to somehow specify hooksPath on their machines too. Ideally when cloning your repo. And of course they shouldn't do it manually because they will forget.

There are ways to automate this process and their complexity will vary from project to project. With Node-based projects it will be easy. Just add a preinstall script in your package.json:
```
{
  ......
  "scripts": {
    "preinstall": "git config core.hooksPath hooks_dir",
    .......
  }
}
```
and now you have a guarantee that everyone gets a proper hooksPath because everyone has to run __npm install__. No extra action will be necessary in such scenario.

HooksPath can also be stored in a .gitconfig file (though it doesn't change things dramatically). This file can be added to version control and shared without difficulty. But it won't work just like that because you have to explicitly __enable__ this config using a command
```
git config --local include.path ../.gitconfig
```
This too can be automated via Node's preinstall or any other tools available in your project.

__Solution 3.__ Some people suggest using symlinks. I'll quote from StackOverflow: "You could create a hooks directory in your project directory with all the scripts, and then symlink them in .git/hooks. Of course, each person who cloned the repo would have to set up these symlinks."
This solution is very similar to #2 (hooksPath) and can be automated in a similar way. Though hooksPath is preferable because it's more "native".

__Solution 4.__ Another approach is to keep your hooks in a separate repo. This approach is suggested by [pre-commit](https://github.com/observing/pre-commit) hook installer (but it's only for pre-commit hooks). Storing hooks in a repo, just like solutions #2 and #3, makes it easy to share hooks but difficult to run them when necessary.
