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
| `v0.5.4` | Blob upload accepts any 2xx, not only 200 (RFC 8620 §6.1; yarilo answers 201) | cherry-pick of `25f7b38d` from the stale GitHub repo | not on sr.ht; the change would arrive there under another sha, so the next tag is watched and checked by hand |
| `v0.5.5` | A /get request sends an empty ids as `[]` and `#ids` without `ids` (RFC 8620 §3.7, §5.1) | ours; dropped when upstream fixes the same | the change would arrive there under another sha, so the next tag is watched and checked by hand |

Upstream takes patches by mail on sr.ht, so a row names the upstream commit or
tag our patch waits for, not a pull request. Every row here has a row in the
watch list of `.github/workflows/sync-upstream.yml`, added in the same change:
a patch in this table and not in the list is a patch nobody is watching.
