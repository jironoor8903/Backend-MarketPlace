# Backend-Marketplace
### Requirement Detail:
#### In the platform, users can view new products for sale, and comment on these products.

* Users need to register and login
* Users can browse a paginated list of products no matter they login or not.(a paginated list contains at most 20 items, and number of
paginated list is no more than 20)
* Users can search for specific products
* Users can view product details, including title, descriptions, product categories, product photos, and list of comments.
* Users can comment (including reply to a user) on a product.

### Performance Requirement:
#### We assume that the rows of the core table(products) is more than 50000000 (each product have at lease two categories and at lease 5 comment with
one sub comment), and the number of user is more than 2000. our system should support the follow performance metrics:
* >= 3000 qps (browse list of products api) and >= 300 tps(login api)
* <= 300ms response time of all the requests

