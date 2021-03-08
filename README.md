# Simple Burger API

## Description

Create API (backend code + db) similar to this one [https://punkapi.com/documentation/v2](https://punkapi.com/documentation/v2), but for Burgers.

**IMPORTANT:** In Addition to the described methods in the documentation you need to provide a method for creating new Burger.

**IMPORTANT 2:** You do not need to create all the PARAMs(abv_gt, abv_lt, etc). Just create one for burger_name.

You can use random burger images from around the internet :)

Whether you deploy the api on public endpoint (like AWS/heroku) or keep it localhost is up to you and bears no significance on the review of your work.

*This description is purposefully short in order to allow you to be creative*

# Getting Started
The v1 API supports endpoints for fetching and creating burgers

## Dependencies
The application requires the following prerequisites for building and running:

Go 1.13 or higher

## Building and running locally
The project can be build and run via:
```bash
go build .
go run ./cmd/burger-api.go -migrate true -config ./config.yaml
```

Parameter options:
 * migrate - Perform mongo collection validation setup & fixture population
 * config  - Set configuration path

##Pagination
Requests that return multiple items will be limited to 5 results by default. You can access other pages using the ?page paramater, you can also increase the amount of beers returned in each request by changing the ?per_page paramater.

```bash
curl http://localhost:8080/api/v1/burgers/?page=2&per_page=80
```

## Get Burgers

Gets a page of burgers (by default page=0, per_page=5)

```bash
curl http://localhost:8080/api/v1/burgers/
```

Optionally, to filter by name, use query param `burger_name`, ex:

```bash
curl http://localhost:8080/api/v1/burgers?burger_name=Ultra-Smashed%20Cheeseburgers%20Recipe
```

## Get a Single Burger

Gets a burger by providing burgers id

```bash
curl http://localhost:8080/api/v1/burgers/604401ebcc64802b255a200a
```

## Get a Random Burger

Gets a random burger from the API, this takes no paramaters.


```bash
curl http://localhost:8080/api/v1/burgers/random
```

## Save a Burger

Send POST request with json body for a burger. Required fields are name and ingredients

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"My Special Burger","ingredients":[ "bread", "mustard"]}' \
  http://localhost:8080/api/v1/burgers
```

## Example response

```json
{
  "id": "60464ed246b44aaba7e97673",
  "name": "Ultra-Smashed Cheeseburgers Recipe",
  "ingredients": [
    "1 soft hamburger roll, buttered and toasted",
    "Condiments and toppings as desired, such as mayonnaise, mustard, shredded lettuce, onions, tomatoes, and pickles",
    "4 ounces (110g) freshly ground beef chuck, divided into two 2-ounce (55g) balls",
    "Kosher salt and freshly ground black pepper",
    "1 slice good melting cheese, such as American, cheddar, or homemade melting cheese"
  ],
  "notes": [
    "Smoky -- cook it outside!"
  ],
  "source": "https://www.seriouseats.com/recipes/2014/03/ultra-smashed-cheeseburger-recipe-food-lab.html",
  "instructions": [
    "Prepare burger bun by laying toppings on bottom half of bun. Have it nearby and ready for when your burger is cooked.",
    "Preheat a large stainless steel sauté pan or skillet over high heat for 2 minutes. Place balls of beef in pan and smash down with a stiff metal spatula, using a second spatula to add pressure. Smashed patties should be slightly wider than burger bun.",
    "Season generously with salt and pepper and allow to cook until patties are well browned and tops are beginning to turn pale pink/gray in spots, about 45 seconds. Using a bench scraper or the back side of a stiff metal spatula, carefully scrape patties from pan, making sure to get all of the browned bits.",
    "Flip patties and immediately place a slice of cheese over 1 patty, then stack the second directly on top. Immediately remove from pan and transfer to waiting burger bun. Serve."
  ],
  "encodedImage": "..."
}
```