export const metadata = {
    title: 'Login Page',
    description: '',
  }

export default function LoginLayout({
    children,
  }: {
    children: React.ReactNode;
  }) {
    return <section>{children}</section>;
  }