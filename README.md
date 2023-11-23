### TODO

- Tailwind & htmx set up properly. not from CDN
  - https://youtu.be/aRCqCZosejc
-

## images

200x200 thumbnail version
processed version at max size w/h x (max image upload size = 1mb?)
compress both images

- https://github.com/h2non/bimg
- https://github.com/disintegration/gift

## Paths

- /
- /signin
- /signup

- /dashboard
  - general info
- /sites
  - show all your sites
- /sites/edit/id
  - show one site edit
- /sites/preview/id
  - view the actual site
- /account
  - set up cloud accounts
  - change password etc

site actions

- view
- download
- edit
- regenerate
- clone
- upload to
- delete

Site data

- target url

favicon

content

- titles
- descriptions
- keywords
- paragraphs
  faq
- file (q & as)
- images, from file or from direct link

  - imgs_uploaded
  - imgs_links
  - or only one but https =url else image id

- iframe/links

- legal pages
- enable privacy, terms, cookies
- customize? -> template

business

- logo
- name
- phone
- email
- website
- adress
- socials

Google embed

relatedPages

options
relatedPages

- chance to appear
- append
- chance to append

faq

- questions per page

img opts

- optional geotag images
- optional randomize image tint & size

## Set up local supabase instance

```bash
pnpm install
# run these commands prefix with pnpm
# https://supabase.com/docs/guides/cli/managing-environments
```

## Check if user is a paid user using

user metadata
supa.auth.admin.UpdateUserById

- doesnt exist on the client I am using :(

Ill just need to make a PR here

- https://github.com/nedpals/supabase-go/blob/master/auth.go
