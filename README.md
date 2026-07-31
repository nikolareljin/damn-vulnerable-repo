# damn-vulnerable-repo

**A deliberately leaky repository, for testing secret scanners.**

> ### ⚠️ Everything in this repository is fake
>
> No real credential has ever been committed here. The values are synthetic strings of
> the right *shape* so that scanners detect them:
>
> - **AWS** uses Amazon's own published example key, `*****`
> - **Stripe** values carry the `sk_test_` prefix, which is a test-mode key by definition
> - **Every other token** is a made-up string matching the provider's format
> - **SSH keys** are generated fresh by the seeding script and used nowhere
>
> Nothing here grants access to anything. Please do not report these as leaked
> credentials — that is the entire point of the repository.

## What this is for

It exists to exercise [**Leak Lock**](https://github.com/nikolareljin/leak-lock), a VS Code
extension that finds secrets in git history and removes them. Testing a scanner needs a
repository that is *reliably* messy in specific ways, and messing up a real one is a poor
idea.

The fixture is generated, not hand-written, by
[`tools/seed-fake-leaks.sh`](https://github.com/nikolareljin/leak-lock/blob/main/tools/seed-fake-leaks.sh)
in the Leak Lock repository. It can be re-run to restore the fixture after a cleanup has
destroyed it, which is the normal testing loop: **seed → scan → clean → seed again**.

## Using it

```bash
git clone https://github.com/nikolareljin/damn-vulnerable-repo.git
cd /path/to/leak-lock

# Seed, with a throwaway local remote so force-push and verification work
bash tools/seed-fake-leaks.sh ../damn-vulnerable-repo --local-remote
```

Then point Leak Lock at the clone and scan. Full documentation:
[docs/TEST_FIXTURE.md](https://github.com/nikolareljin/leak-lock/blob/main/docs/TEST_FIXTURE.md).

Re-seed after a cleanup:

```bash
bash tools/seed-fake-leaks.sh ../damn-vulnerable-repo --reset --push
```

## What the fixture contains

Everything lands under `leaklock-fixture/` on branches prefixed `leaklock-fixture/`, so it
never mixes with anything else.

| Branch | Contents |
|---|---|
| `leaklock-fixture/main-leaks` | The bulk of the history |
| `leaklock-fixture/hotfix-db-creds` | MySQL and MongoDB URLs with passwords |
| `leaklock-fixture/release-1.0` | Azure, SendGrid and Twilio credentials |
| `leaklock-fixture/legacy-import` | `htpasswd` and a Docker auth blob, added then deleted |
| `leaklock-fixture/dev-alice` | Fine-grained GitHub token, Slack bot token, JWT |
| `leaklock-fixture/experimental` | A Stripe test key, unreachable from the others |
| `leaklock-fixture/ops-remote-only` | Exists **only** on the remote |

Credential types: AWS key and secret · GitHub PAT, classic and fine-grained · Slack bot
token and webhook · Stripe test keys · Google API key · SendGrid · Twilio · npm and PyPI
tokens · JWT · Azure storage connection string · Postgres, MySQL and MongoDB URLs with
passwords · `.netrc` and `.htpasswd` · Docker auth blob · GCP service-account JSON · RSA
and ed25519 private keys.

### Cases it is built to cover

| Path | Exercises |
|---|---|
| `config/settings.py` | One secret across **two commits**, still on disk |
| `config/deploy_key*` | Private keys in **history only** — deleted later |
| `ci/deploy.sh` | One token matching **two rules** |
| `node_modules/@vendor/` | A finding inside a third-party dependency |
| `env`, `secrets/` | **Untracked** — the fix is deleting the file, not rewriting history |
| `big-blob.txt` | Over a scanner's default size limit |
| `ops-remote-only` | Found only if refs are refreshed before scanning |
| `docs/runbook.md` | An internal hostname and a customer identifier that **no scanner flags** |

## `main` is intentionally empty

`main` holds only this README. The fixture lives entirely on the
`leaklock-fixture/*` branches, so cleaning them up — including force-pushing rewritten
history — never touches the branch that explains what this repository is.

## Licence

MIT. Do what you like with it.
