// Store.ts
export interface Store {
    id: string;
    name: string;
    userId: string;
    createdAt: Date;
    updatedAt: Date;
}

// Size.ts
export interface Size {
    id: string;
    storeId: string;
    name: string;
    value: string;
    createdAt: Date; // Consider changing to Date if appropriate
    updatedAt: Date; // Consider changing to Date if appropriate
}

// Product.ts
export interface Product {
    id: string;
    storeId: string;
    categoryId: string;
    name: string;
    price: number;
    isFeatured: boolean;
    isArchived: boolean;
    sizeId: string;
    colorId: string;
    createdAt: Date;
    updatedAt: Date;
}

// Order.ts
export interface Order {
    id: string;
    storeId: string;
    isPaid: boolean;
    phone: string;
    address: string;
    createdAt: Date;
    updatedAt: Date;
}

// OrderItem.ts
export interface OrderItem {
    id: string;
    orderId: string;
    productId: string;
}

// Image.ts
export interface Image {
    id: string;
    productId: string;
    url: string;
    createdAt: Date;
    updatedAt: Date;
}

// Color.ts
export interface Color {
    id: string;
    storeId: string;
    name: string;
    value: string;
    createdAt: Date;
    updatedAt: Date;
}

// Category.ts
export interface Category {
    id: string;
    storeId: string;
    store: Store;
    billboardId: string;
    name: string;
    createdAt: Date;
    updatedAt: Date;
}

// Billboard.ts
export interface Billboard {
    id: string;
    storeId: string;
    label: string;
    isActive: boolean;
    imageUrl: string;
    createdAt: Date;
    updatedAt: Date;
}
