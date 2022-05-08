# xyzzyapi

## Testable stack

1. Bring up Senzing stack: https://github.com/Senzing/knowledge-base/blob/main/HOWTO/test-with-latest.md

1. Set environment variables.

    ```console
    export SENZING_G2_DIR=~/senzing-3.0.0
    ```

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=docktermj
    export GIT_REPOSITORY=xyzzygoapi
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"
    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.
1. Checkout branch.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    git checkout issue.1.dockter.2
    ```

1. Set environment variables.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make -e test
    git checkout issue.1.dockter.2
    ```
