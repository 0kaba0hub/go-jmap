# yarilo-patches branch

This file lives **only on the `yarilo-patches` branch** of the
[0kaba0hub/go-jmap](https://github.com/0kaba0hub/go-jmap) fork. The `main`
branch mirrors [git.sr.ht/~rockorager/go-jmap](https://git.sr.ht/~rockorager/go-jmap)
verbatim and must never carry downstream changes.

The upstream is the module's own repository on sr.ht, which the Go module proxy
serves. The GitHub repository `rockorager/go-jmap` stopped in 2023 and carries a
commit the module does not, so it is not mirrored.

## Why this fork exists

[yarilo](https://github.com/yarilomail/yarilo) uses the module in tests and the
smoketest probe of JMAP. Each change we need before upstream has it is a
cherry-pick on top of `main`; when upstream takes it, the cherry-pick is
dropped. With no patches, yarilo takes the upstream tag directly and has no
`replace`; the `replace` onto this fork arrives with the first patch.

| Upstream change awaited (commit or tag) | Topic | Our commit | State |
|:---|:---|:---|:---|

Upstream takes patches by mail on sr.ht, so a row names the upstream commit or
tag our patch waits for, not a pull request. Every row here has a row in the
watch list of `.github/workflows/sync-upstream.yml`, added in the same change:
a patch in this table and not in the list is a patch nobody is watching.
