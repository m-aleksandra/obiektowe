import Fluent
import Vapor

struct ProductController: RouteCollection {
    func boot(routes: any RoutesBuilder) throws {
        let products = routes.grouped("products")

        // JSON API
        products.get(use: self.index)              // GET /products (JSON)
        products.post(use: self.create)            // POST /products (JSON)
        products.group(":productID") { product in
            product.delete(use: self.delete)       // DELETE /products/:id (JSON)
            product.post("edit", use: self.editForm)      // POST /products/:id/edit (Form)
            product.post("delete", use: self.deleteForm)  // POST /products/:id/delete (Form)
        }

        // Form HTML
        products.post("create", use: self.createForm)      // POST /products/create (Form)
    }

    // MARK: - JSON API

    @Sendable
    func index(req: Request) async throws -> [ProductDTO] {
        try await Product.query(on: req.db).all().map { $0.toDTO() }
    }

    @Sendable
    func create(req: Request) async throws -> ProductDTO {
        let dto = try req.content.decode(ProductDTO.self)
        let product = dto.toModel()
        try await product.save(on: req.db)
        return product.toDTO()
    }

    @Sendable
    func delete(req: Request) async throws -> HTTPStatus {
        guard let product = try await Product.find(req.parameters.get("productID"), on: req.db) else {
            throw Abort(.notFound)
        }
        try await product.delete(on: req.db)
        return .noContent
    }

    // MARK: - Form handlers (HTML)

    @Sendable
    func createForm(req: Request) async throws -> Response {
        let product = try req.content.decode(Product.self)
        try await product.save(on: req.db)
        return req.redirect(to: "/")
    }

    @Sendable
    func editForm(req: Request) async throws -> Response {
        guard let id = req.parameters.get("productID"),
              let uuid = UUID(uuidString: id),
              let existing = try await Product.find(uuid, on: req.db)
        else {
            throw Abort(.notFound)
        }

        let updated = try req.content.decode(Product.self)
        existing.name = updated.name
        existing.description = updated.description
        existing.price = updated.price
        try await existing.save(on: req.db)
        return req.redirect(to: "/")
    }

    @Sendable
    func deleteForm(req: Request) async throws -> Response {
        guard let product = try await Product.find(req.parameters.get("productID"), on: req.db) else {
            throw Abort(.notFound)
        }

        try await product.delete(on: req.db)
        return req.redirect(to: "/")
    }
}
