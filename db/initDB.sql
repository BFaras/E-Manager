CREATE SCHEMA IF NOT EXISTS public;

-- CreateTable: Store
CREATE TABLE public."Store" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "name" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- CreateTable: Billboard
CREATE TABLE public."Billboard" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "storeId" TEXT NOT NULL,
    "label" TEXT NOT NULL,
    "imageUrl" TEXT NOT NULL,
    "isActive" BOOLEAN DEFAULT false NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Billboard_storeId_fkey" FOREIGN KEY ("storeId") REFERENCES public."Store"("id") ON DELETE CASCADE
);

-- CreateTable: Category
CREATE TABLE public."Category" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "storeId" TEXT NOT NULL,
    "billboardId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Category_storeId_fkey" FOREIGN KEY ("storeId") REFERENCES public."Store"("id") ON DELETE CASCADE,
    CONSTRAINT "Category_billboardId_fkey" FOREIGN KEY ("billboardId") REFERENCES public."Billboard"("id") ON DELETE CASCADE
);

-- CreateTable: Size
CREATE TABLE public."Size" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "storeId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "value" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Size_storeId_fkey" FOREIGN KEY ("storeId") REFERENCES public."Store"("id") ON DELETE CASCADE
);

-- CreateTable: Color
CREATE TABLE public."Color" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "storeId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "value" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Color_storeId_fkey" FOREIGN KEY ("storeId") REFERENCES public."Store"("id") ON DELETE CASCADE
);

-- CreateTable: Product
CREATE TABLE public."Product" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "storeId" TEXT NOT NULL,
    "categoryId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "price" DECIMAL(10,2) NOT NULL,
    "count" INTEGER DEFAULT 0 NOT NULL,
    "isFeatured" BOOLEAN DEFAULT false NOT NULL,
    "isArchived" BOOLEAN DEFAULT false NOT NULL,
    "isDeleted" BOOLEAN DEFAULT false NOT NULL,
    "sizeId" TEXT NOT NULL,
    "colorId" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Product_storeId_fkey" FOREIGN KEY ("storeId") REFERENCES public."Store"("id") ON DELETE CASCADE,
    CONSTRAINT "Product_categoryId_fkey" FOREIGN KEY ("categoryId") REFERENCES public."Category"("id") ON DELETE CASCADE,
    CONSTRAINT "Product_sizeId_fkey" FOREIGN KEY ("sizeId") REFERENCES public."Size"("id") ON DELETE CASCADE,
    CONSTRAINT "Product_colorId_fkey" FOREIGN KEY ("colorId") REFERENCES public."Color"("id") ON DELETE CASCADE
);

-- CreateTable: Image
CREATE TABLE public."Image" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "productId" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Image_productId_fkey" FOREIGN KEY ("productId") REFERENCES public."Product"("id") ON DELETE CASCADE
);

-- CreateTable: Order
CREATE TABLE public."Order" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "storeId" TEXT NOT NULL,
    "isPaid" BOOLEAN DEFAULT false NOT NULL,
    "phone" TEXT NOT NULL DEFAULT '',
    "address" TEXT NOT NULL DEFAULT '',
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT "Order_storeId_fkey" FOREIGN KEY ("storeId") REFERENCES public."Store"("id") ON DELETE CASCADE
);

-- CreateTable: OrderItem
CREATE TABLE public."OrderItem" (
    "id" TEXT DEFAULT gen_random_uuid()::TEXT PRIMARY KEY,
    "orderId" TEXT NOT NULL,
    "productId" TEXT NOT NULL,
    CONSTRAINT "OrderItem_orderId_fkey" FOREIGN KEY ("orderId") REFERENCES public."Order"("id") ON DELETE CASCADE,
    CONSTRAINT "OrderItem_productId_fkey" FOREIGN KEY ("productId") REFERENCES public."Product"("id") ON DELETE CASCADE
);
