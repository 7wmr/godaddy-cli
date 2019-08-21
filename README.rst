GoDaddy Command Line Tool
==========================

.. image:: https://circleci.com/gh/7wmr/godaddy-cli.svg?style=shield
    :target: https://circleci.com/gh/7wmr/godaddy-cli

Develop
--------------------------

Clone repository.

.. code:: bash

  git clone git@github.com:7wmr/godaddy-cli.git ~/go/src/github.com/7wmr/godaddy-cli -b develop

Link repository to repo directory.

.. code:: bash

  ln -s ~/go/src/github.com/7wmr/godaddy-cli ~/repos/godaddy-cli

Generate command documentation.

.. code:: bash

   go run docs.go


Deploy
--------------------------

.. code:: bash

  brew tap 7wmr/homebrew-tools


.. code:: bash

  brew install godaddy-cli


Usage
--------------------------

- `godaddy <docs/godaddy.rst>`_
- `godaddy domain <docs/godaddy_domain.rst>`_
- `godaddy domain set-record <docs/godaddy_domain_set-record.rst>`_
