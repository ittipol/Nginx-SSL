export const metadata = {
    title: 'Register Page',
    description: '',
  }

export default function RegisterLayout({
    children,
  }: {
    children: React.ReactNode;
  }) {
    return <section>{children}</section>;
  }