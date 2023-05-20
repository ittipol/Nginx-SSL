export const metadata = {
    title: 'Profile',
    description: '',
  }

export default function UserProfileLayout({
    children,
  }: {
    children: React.ReactNode;
  }) {
    return <section>{children}</section>;
  }