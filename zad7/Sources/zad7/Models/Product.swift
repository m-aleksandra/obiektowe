import Fluent
import struct Foundation.UUID

final class Product: Model, @unchecked Sendable {
    static let schema = "products"

    @ID(key: .id)
    var id: UUID?

    @Field(key: "name")
    var name: String

    @Field(key: "description")
    var description: String

    @Field(key: "price")
    var price: Double

    init() {}

    init(id: UUID? = nil, name: String, description: String, price: Double) {
        self.id = id
        self.name = name
        self.description = description
        self.price = price
    }

    func toDTO() -> ProductDTO {
        .init(
            id: self.id,
            name: self.$name.value,
            description: self.$description.value,
            price: self.$price.value
        )
    }
}
