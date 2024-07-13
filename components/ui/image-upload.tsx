"use client";

interface ImageUploadProps {
    disabled?: boolean;
    onChange: (value: string) => void;
    onRemove: (value: string) => void;
    value: string[];
}   

import React from 'react'


export default function ImageUpload({
    disabled,
    onChange,
    onRemove,
    value,
 }: ImageUploadProps
) {

    const [isMounted, setIsMounted] = useState(false);

    useEffect(() => {
        setIsMounted(true);
    }, [])

    if (!isMounted) {
        return null;
    }

  return (
    <div>
      ImageUpload
    </div>
  )
}

