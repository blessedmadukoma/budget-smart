import { message, notification } from "ant-design-vue";
import type { NotificationArgsProps } from "ant-design-vue/es/notification";

export function useNotification() {
  const success = (content: string, title = "Success", duration = 4.5) => {
    notification.success({
      message: title,
      description: content,
      duration,
    });
  };

  const error = (content: string, title = "Error", duration = 4.5) => {
    notification.error({
      message: title,
      description: content,
      duration,
    });
  };

  const warning = (content: string, title = "Warning", duration = 4.5) => {
    notification.warning({
      message: title,
      description: content,
      duration,
    });
  };

  const info = (content: string, title = "Information", duration = 4.5) => {
    notification.info({
      message: title,
      description: content,
      duration,
    });
  };

  const toast = {
    success: (content: string, duration = 2) =>
      message.success(content, duration),
    error: (content: string, duration = 2) => message.error(content, duration),
    warning: (content: string, duration = 2) =>
      message.warning(content, duration),
    info: (content: string, duration = 2) => message.info(content, duration),
    loading: (content: string, duration = 2) =>
      message.loading(content, duration),
  };

  const custom = (options: NotificationArgsProps) => {
    notification.open(options);
  };

  return {
    success,
    error,
    warning,
    info,
    toast,
    custom,
    notification,
    message,
  };
}
