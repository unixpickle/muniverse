# Building a container

To build the container, you must first [download the games](../games). Next, copy the downloaded games to this directory:

```
cp -r ../games/downloaded_games .
```

Now you can run `docker build`:

```
docker build --tag muniverse:VERSION .
```
