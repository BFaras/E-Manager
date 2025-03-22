
export interface Store {
    id: string;
    name: string;
    userId: string;
    createdAt: Date;
    updatedAt: Date;
}

export interface Size {
    id: string;
    storeId: string;
    name: string;
    value: string;
    createdAt: Date; 
    updatedAt: Date; 
}

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

export interface Order {
    id: string;
    storeId: string;
    isPaid: boolean;
    phone: string;
    address: string;
    createdAt: Date;
    updatedAt: Date;
}

export interface OrderItem {
    id: string;
    orderId: string;
    productId: string;
}

export interface Image {
    id: string;
    productId: string;
    url: string;
    createdAt: Date;
    updatedAt: Date;
}

export interface Color {
    id: string;
    storeId: string;
    name: string;
    value: string;
    createdAt: Date;
    updatedAt: Date;
}

export interface Category {
    id: string;
    storeId: string;
    store: Store;
    billboardId: string;
    name: string;
    createdAt: Date;
    updatedAt: Date;
}

export interface Billboard {
    id: string;
    storeId: string;
    label: string;
    isActive: boolean;
    imageUrl: string;
    createdAt: Date;
    updatedAt: Date;
}
