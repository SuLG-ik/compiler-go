import './StatusBar.css'

interface StatusBarProps {
  status: string
  row: number
  col: number
}

export function StatusBar({ status, row, col }: StatusBarProps) {
  return (
    <div className="statusbar">
      <span className="statusbar__status">{status}</span>
      <span className="statusbar__cursor">Стр. {row}  Стб. {col}</span>
    </div>
  )
}
