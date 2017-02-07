lister-gen is a tool to generate listers for a client.

lister-gen requires the `// +genclient=true` annotation on each
exported type in both the unversioned base `types.go` as well as each
specifically versioned `types.go`.

lister-gen requires the `// +groupName=` annotation on the `doc.go` in
both the unversioned base directory as well as in each specifically
versioned directory. The annotation does not have to have any specific
content, but it does have to exist.
