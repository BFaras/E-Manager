"use client";

import React, { useCallback } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from './dialog';

interface ModalProps {
    title: string;
    description: string;
    isOpen: boolean;
    onClose: () => void;
    children?: React.ReactNode;
}

export function Modal(props: ModalProps) {

    const onChangeHandle = useCallback((open: boolean) => {
        if (!open) {
            props.onClose();
        }
    }, []);

    return (
        <Dialog open={props.isOpen} onOpenChange={onChangeHandle}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>{props.title}</DialogTitle>
                    <DialogDescription>{props.description}</DialogDescription>
                </DialogHeader>
                <div>{props.children}</div>
            </DialogContent>
        </Dialog>
    );
}

export default Modal;