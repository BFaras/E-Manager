"use client";
import React, { useEffect, useState } from "react";
import Heading from "@/components/heading";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Trash } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormField, FormItem, FormLabel } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useParams, useRouter } from "next/navigation";
import axios from "axios";
import toast from "react-hot-toast";
import { AlertModal } from "@/app/modals/alert-modal";
import { ApiAlert } from "@/components/ui/api-alert";
import { useOrigin } from "@/hooks/use-origin";
import { Store } from "@prisma/client";
import { useAuth } from "@clerk/nextjs";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

interface SettingFormProps {
  initialData: Store;
}

const formSchema = z.object({
  name: z.string().min(1),
});

type SettingFormValues = z.infer<typeof formSchema>;

export default function SettingForm({ initialData }: SettingFormProps) {
  const params = useParams();
  const router = useRouter();
  const origin = useOrigin();
  const {getToken} = useAuth()


  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const form = useForm<SettingFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: initialData,
  });

  const setup = async () => {
    await setUpInterceptor(getToken);
  };

  useEffect(() => {
    setup();
  }, [getToken]);

  const onSubmit = async (data: SettingFormValues) => {
    try {
      setLoading(true);
      await axiosInstance.patch(`secured/stores/${params.storeId}`, data);
      router.refresh();
      setLoading(false);
      toast.success("store updated");
    } catch (error) {
      toast.error("Something went wrong");
    } finally {
      setLoading(false);
      setOpen(false);
    }
  };

  const onDelette = async () => {
    try {
      setLoading(true);
      await axiosInstance.delete(`secured/stores/${params.storeId}`);
      router.refresh();
      router.push("/");
      toast.success("Store deleted");
    } catch (error) {
      toast.error("Make sure you removed all the products first");
    } finally {
      setLoading(false);
      setOpen(false);
    }
  };
  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={onDelette}
        loading={loading}
      ></AlertModal>
      <div className="flex items-center justify-between">
        <Heading title="Settings" description="Manage store preference" />
        <Button variant="destructive" size="icon" onClick={() => setOpen(true)}>
          <Trash className="h-4 w-4"></Trash>
        </Button>
      </div>
      <Separator></Separator>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8 w-full "
        >
          <div className="grid grid-cols-3 gap-8">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
                      placeholder="Store name"
                      {...field}
                    ></Input>
                  </FormControl>
                </FormItem>
              )}
            ></FormField>
          </div>
          <Button disabled={loading} className="ml-auto" type="submit">
            Save Changes
          </Button>
        </form>
      </Form>
      <Separator></Separator>
      <ApiAlert
        title="NEXT_PUBLIC_API_URL"
        description={`${origin}/api/${params.storeId}`}
        variant="public"
      ></ApiAlert>
    </>
  );
}
