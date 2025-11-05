# Decklist to Scryfall

Converts a text deck list into a URL you can use in Scryfall to look up all the cards in your deck at once and see all printings of all cards.

Output will be broken up into chunks of 20 cards since the Scryfall query doesnt work well with tons of cards at once

## TODO:

- filter out Basic lands, add filter args for common artifacts (Command tower, Sol ring, Arcan signet, etc)

# Usage

```bash
$ go run main.go examples/Hakbal.txt

Cards 1-20:
https://scryfall.com/search?as=grid&order=name&q=%22Aetherize%22+or+%22Alchemist%27s+Refuge%22+or+%22Arcane+Signet%22+or+%22Beast+Within%22+or+%22Benthic+Biomancer%22+or+%22Branching+Evolution%22+or+%22Breeding+Pool%22+or+%22Cold-Eyed+Selkie%22+or+%22Command+Tower%22+or+%22Cyclonic+Rift%22+or+%22Deeproot+Elite%22+or+%22Deeproot+Historian%22+or+%22Deeproot+Pilgrimage%22+or+%22Deeproot+Waters%22+or+%22Dreamroot+Cascade%22+or+%22Emperor+Mihail+II%22+or+%22Explore%22+or+%22Fierce+Guardianship%22+or+%22Flooded+Grove%22+or+%22Forerunner+of+the+Heralds%22+%28game%3Apaper%29+unique%3Aprints

Cards 21-40:
https://scryfall.com/search?as=grid&order=name&q=%22Forest%22+or+%22Growth+Spiral%22+or+%22Harbinger+of+the+Seas%22+or+%22Hardened+Scales%22+or+%22Herald+of+Secret+Streams%22+or+%22Heroic+Intervention%22+or+%22Hinterland+Harbor%22+or+%22Inspiring+Call%22+or+%22Island%22+or+%22Karn%27s+Bastion%22+or+%22Kindred+Discovery%22+or+%22Kiora%27s+Follower%22+or+%22Kodama%27s+Reach%22+or+%22Kopala%2C+Warden+of+Waves%22+or+%22Kumena%2C+Tyrant+of+Orazca%22+or+%22Lord+of+Atlantis%22+or+%22Master+of+the+Pearl+Trident%22+or+%22Merfolk+Cave-Diver%22+or+%22Merfolk+Mistbinder%22+or+%22Merfolk+Skydiver%22+%28game%3Apaper%29+unique%3Aprints

Cards 41-60:
https://scryfall.com/search?as=grid&order=name&q=%22Merfolk+Sovereign%22+or+%22Merrow+Commerce%22+or+%22Merrow+Reejerey%22+or+%22Metallic+Mimic%22+or+%22Mist+Dancer%22+or+%22Misty+Rainforest%22+or+%22Nature%27s+Lore%22+or+%22Nicanzil%2C+Current+Conductor%22+or+%22Ozolith%2C+the+Shattered+Spire%22+or+%22Path+of+Ancestry%22+or+%22Raise+the+Palisade%22+or+%22Rapid+Hybridization%22+or+%22Realmwalker%22+or+%22Reflections+of+Littjara%22+or+%22Rejuvenating+Springs%22+or+%22Reliquary+Tower%22+or+%22Rhystic+Study%22+or+%22Ripples+of+Potential%22+or+%22Roaming+Throne%22+or+%22Rogue%27s+Passage%22+%28game%3Apaper%29+unique%3Aprints

Cards 61-80:
https://scryfall.com/search?as=grid&order=name&q=%22Seafloor+Oracle%22+or+%22Secluded+Courtyard%22+or+%22Simic+Ascendancy%22+or+%22Simic+Growth+Chamber%22+or+%22Simic+Signet%22+or+%22Singer+of+Swift+Rivers%22+or+%22Sol+Ring%22+or+%22Stonybrook+Banneret%22+or+%22Svyelun+of+Sea+and+Sky%22+or+%22Swiftfoot+Boots%22+or+%22Tatyova%2C+Benthic+Druid%22+or+%22Temple+of+Mystery%22+or+%22The+Ozolith%22+or+%22Thieving+Skydiver%22+or+%22Tishana%27s+Tidebinder%22+or+%22Topography+Tracker%22+or+%22Unclaimed+Territory%22+or+%22Vineglimmer+Snarl%22+or+%22Vodalian+Hexcatcher%22+or+%22Wanderwine+Prophets%22+%28game%3Apaper%29+unique%3Aprints

Cards 81-85:
https://scryfall.com/search?as=grid&order=name&q=%22Wave+Goodbye%22+or+%22Wild+Rose+Rebellion%22+or+%22Yavimaya+Coast%22+or+%22Zegana%2C+Utopian+Speaker%22+or+%22Hakbal+of+the+Surging+Soul%22+%28game%3Apaper%29+unique%3Aprints

```

# Resources

- https://scryfall.com/docs/syntax#exact , https://scryfall.com/advanced