##Option1:

To add v126

```
copy v125/ to v126/
edit v126/ interface
edit client/ interface & switch
edit v123/, v124/ & v125/ to comply to the new interface
```

Pros: easy 

Cons:
* you have to update all previous versions when you update the client interface, so we will probably support only X last versions
* lots of code duplication - a bug or fix that impacts multiple versions requires changes to many duplicate files


---

## Option2:

To add v126

```
create v126/ implementation and adaptor to v125
update main.go to use v126.NewClient()
```

Pros: no need to touch the previous version when you add one, so we could potentially support all previous version

Cons:
* more work to create a new version
* it is difficult to see the differences between versions
* a bug or fix that impacts multiple versions requires changes to many duplicate files


## Option3:

Put all new flags and endpoints behind "version gates" that report a friendly 
error to the user when they use a version that doesn't support that feature.

No extra code is required to add a new version. Flags are added inline with the 
PR to add the feature.


Pros:
* it is easy for contributors to know where to add code
* no code duplicate for every version
* it is much easier to see the diff between versions (`git grep 'if ... vx'` or 
  `git diff v1..v2`)
* it should be easier to support "experimental as runtime flag" because experimental
  features can be gated the same way

Cons:
* if a field changes type between versions it would require a custom type which 
  can handle the variation during serialization/deserialization
