# Development roadmap for talker

The following items are what I will strive to have in v1.
Any pull request, solution, thought or draft on these topics would be great. Feel free to share them in issues.

* [x] Migration to a fully golang runtime managed API with **no CGO (gcc, g++) dependencies**.
	* this will reduce our dependency on 3rd party modules.
	* things get easier later when we want to add MAC or linux support.
	* The issue of UTF-8/16 can be challenging but I think it's worth it.
* [ ] Support for Linux speech apis
* [ ] Support for Darwin's speech apis
	* We can solve these two features with golang's tags.



