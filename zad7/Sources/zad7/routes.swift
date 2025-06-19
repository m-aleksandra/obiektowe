import Fluent
import Vapor

func routes(_ app: Application) throws {
    app.get { req -> EventLoopFuture<View> in
        Product.query(on: req.db).all().flatMap { products in
            req.view.render("index", ["products": products])
        }
    }

    try app.register(collection: ProductController())
}

