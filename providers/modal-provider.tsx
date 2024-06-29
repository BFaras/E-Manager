"use client";


import { StoreModal } from '@/app/modals/store-modal';
import React, { useEffect, useState } from 'react'

export function ModalProvider() {

    const [isMounted, setIsMounted] = useState(false);

    useEffect(() => {
        setIsMounted(true);
    }, [])

    if (!isMounted) {
        return null;
    }

    return (
        <>
            <StoreModal></StoreModal>
        </>
    )

}

