Unicore Network
===============

|travis|_

Universal Core network server.

::

    $ go get -v ./...
    $ go test -v ./...


.. |travis| image:: https://travis-ci.org/praekelt/unicore-network.svg?branch=develop
.. _travis: https://travis-ci.org/praekelt/unicore-network

API
===

::
    GET /identity

Returns the server's identity JSON, generally this document would look
something like this::

    {
        "signature": "uuid4",
        "hostname": "a fqdn",
        "display_name": "a human friendly name, not necessarily unique"
    }

::
    GET /network

Returns a list of all known identities (other than the server's own) that
this specific node is aware of::

    [{
        "signature": "uuid4",
        "hostname": "a fqdn",
        "display_name": "a human friendly name, not necessarily unique"
    }]

::
    PUT /network/:signature

Allows for adding of another nodes identity. The payload must be a JSON
identity document and the document's signature must match the signature
in the URLs path.

::
    GET /network/:signature

Retrieve an identity that was previously put there. If found it returns
the identity document, otherwise

::

    DELETE /network/:signature

Delete an identity. If successful it will return an HTTP 200 OK with the
deleted identity in the response body.
