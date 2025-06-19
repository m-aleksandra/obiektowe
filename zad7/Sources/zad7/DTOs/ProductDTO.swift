import Fluent
import Vapor

struct ProductDTO: Content {
    var id: UUID?
    var name: String?
    var description: String?
    var price: Double?

    func toModel() -> Product {
        let model = Product()
        model.id = self.id

        if let name = self.name {
            model.name = name
        }

        if let description = self.description {
            model.description = description
        }

        if let price = self.price {
            model.price = price
        }

        return model
    }
}
