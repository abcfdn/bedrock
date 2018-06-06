# bedrock
A bedrock protocol for ABC Blockchain Community.


### Good Coding Practice

It is essential for our success to keep our coding standard high. Please only commit clean and well-tested code. There is always space for improvement through code refactoring.

- Unit test is mandatory

- If you make assumptions (for instance, on APIs), test your assumptions

- Pre-commit hook

Install pre-commit hook
```
cd .git/hooks/
ln -sf ../../pre-commit.sh ./pre-commit
```
Add common tests into run_tests.sh for pre-commit checks. In case you have to, you can skip it by using the '--no-verify' option.
```
git commit --no-verify
```
